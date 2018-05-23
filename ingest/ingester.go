package ingest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/juju/ratelimit"
	"github.com/smallnest/gomemcache/memcache"
	"github.com/weibocom/ipc/ingest/weibo"
)

var (
	config *Config

	client = &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 5 * time.Second,
			}).Dial,
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
		},
	}
)

func SetConfig(c *Config) {
	config = c
}

type LongText struct {
	LongTextContent string   `json:"longTextContent"`
	PicIds          []string `json:"pic_ids"`
}

type retryPost struct {
	retries int
	uid     uint64
	mid     uint64
	last    time.Time
}

type post struct {
	uid uint64
	mid uint64
}

// NewConfigWatcher returns a new ConfigWatcher.
func NewConfigWatcher(url string, port int, key string, done chan struct{}, limits int, rate *ratelimit.Bucket) *ConfigWatcher {
	return &ConfigWatcher{
		url:       url,
		port:      port,
		key:       key,
		done:      done,
		ingesters: make(map[string]*Ingester),
		limits:    limits,
		rate:      rate,
	}
}

// ConfigWatcher watches config.properties and start or close ingester.
type ConfigWatcher struct {
	url       string
	port      int
	key       string
	done      chan struct{}
	limits    int
	rate      *ratelimit.Bucket
	ingesters map[string]*Ingester
}

// Watch watches config.properties.
func (u *ConfigWatcher) Watch() {
	for {
		select {
		case <-u.done:
			log.Printf("watcher watch stopped\n")
			return
		default:
			ips, err := net.LookupHost(u.url)
			if err == nil {
				// 确认url是否有新增的ip
				for _, ip := range ips {
					if _, ok := u.ingesters[ip]; !ok {
						log.Printf("a new ip %s resolved out of %s\n", ip, u.url)
						ingester := NewIngester(ip, u.port, u.key, u.done, u.limits, u.rate)
						ingester.start()
						u.ingesters[ip] = ingester
					}
				}
				// 确认是否有老的ip从url里面删除
				if len(ips) > 0 {
					for ip, d := range u.ingesters {
						found := false
						for _, rip := range ips {
							if ip == rip {
								found = true
								break
							}
						}
						if !found {
							d.close()
							delete(u.ingesters, ip)
						}
					}
				}
			}
		}
		time.Sleep(5 * time.Second)
	}
}

// NewIngester returns a neww Ingster.
func NewIngester(ip string, port int, key string, done chan struct{}, limits int, rate *ratelimit.Bucket) *Ingester {
	return &Ingester{
		ip:        ip,
		port:      port,
		key:       key,
		limits:    limits,
		rate:      rate,
		done:      done,
		handlerCh: make(chan *post, config.ChannelBuffer),
		retryCh:   make(chan *retryPost, config.ChannelBuffer),
	}
}

// Ingester reads messages from a given broker.
type Ingester struct {
	ip     string
	port   int
	key    string
	limits int
	rate   *ratelimit.Bucket

	handlerCh chan *post
	retryCh   chan *retryPost
	done      chan struct{}
}

func (d *Ingester) close() {
	log.Printf("drain buffered messages")
	done := make(chan struct{})
	go d.drain(done)
	select {
	case <-done:
		log.Println("finished to drain")
	case <-time.After(time.Minute):
		log.Printf("abandoned draining because it took more than one minute")
	}

	log.Printf("ingester stop reading. ip:%s, port:%d, key:%s\n", d.ip, d.port, d.key)
}

func (d *Ingester) start() {
	go d.handleMessage()
	go d.readMessage()
	go d.retry()
}
func (d *Ingester) readMessage() {
	client := memcache.New(fmt.Sprintf("%s:%d", d.ip, d.port))
	client.Timeout = 30 * time.Second

	rl := ratelimit.NewBucketWithQuantum(time.Second, int64(d.limits), int64(d.limits))
	for {
		select {
		case <-d.done:
			return
		default:
			rl.Wait(1)
			item, err := client.Get(d.key)
			if err != nil {
				log.Printf("failed to get new messages: %v", err)
				continue
			}

			batch := &weibo.TriggerMessageBatch{}
			err = proto.Unmarshal(item.Value, batch)
			if err != nil {
				log.Printf("failed to unmarshal trigger message: %v", err)
				continue
			}

			for _, msg := range batch.Messages {

				body := &weibo.TriggerMessageBody{}
				err := proto.Unmarshal(msg.GetBodyBytes(), body)
				if err != nil {
					log.Printf("failed to unmarshal trigger message: %v", err)
					continue
				}

				status := &weibo.Status{}
				err = proto.Unmarshal(body.GetBody(), status)
				if err != nil {
					log.Printf("failed to unmarshal trigger message: %v", err)
					continue
				}

				uid := status.GetAuthor().GetId()
				mid := status.GetMid()
				level := status.GetAuthor().GetLevel()
				sign := status.GetAuthor().GetSign()

				if level == 2 {
					vflag := (sign >> 6) & ((1 << 4) - 1)
					if vflag >= 1 && vflag <= 7 {

						d.handlerCh <- &post{uid: uid, mid: mid}
					}
				}
			}

		}
	}
}

func (d *Ingester) handleMessage() {
	for {
		select {
		case <-d.done:
			return
		case p := <-d.handlerCh:
			d.rate.Wait(1)
			err := postLongText(p.uid, p.mid)
			if err != nil {
				log.Printf("failed to fetch longtext, uid: %d, mid: %d, err: %v", p.uid, p.mid, err)
				select {
				case d.retryCh <- &retryPost{
					retries: config.Retries,
					uid:     p.uid,
					mid:     p.mid,
					last:    time.Now(),
				}:
				default:
					log.Printf("retry channel if full so {uid: %d, mid: %d} is dropped", p.uid, p.mid)
				}

				continue
			}
		}
	}
}

func postLongText(uid, mid uint64) error {
	u := config.FetchLongTextUrl + "&mids=" + strconv.FormatUint(mid, 10)
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", config.LongTextAuth)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err == nil && len(buf) > 0 {
		m := make(map[string]LongText)
		err = json.Unmarshal(buf, &m)
		if err != nil {
			return err
		}

		v := m[strconv.FormatUint(mid, 10)]

		if v.LongTextContent == "" {
			log.Printf("content is empty for {uid: %d, mid: %d}", uid, mid)
			return nil
		}

		// u = config.PostURL + fmt.Sprintf(`?action=write&company=wb&uid=%d&mid=%d&title=title%d&content=%s`,
		// uid, mid, mid, url.QueryEscape(v.LongTextContent))

		var data = url.Values(make(map[string][]string))
		data.Add("company", "wb")
		data.Add("uid", strconv.FormatUint(uid, 10))
		data.Add("mid", strconv.FormatUint(mid, 10))
		data.Add("title", strconv.FormatUint(mid, 10))
		data.Add("content", v.LongTextContent)

		resp, err := client.PostForm(config.PostURL, data)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		ioutil.ReadAll(resp.Body)

		if resp.StatusCode != 200 {
			return fmt.Errorf("failed to add post: status code: %d", resp.StatusCode)
		}

		log.Printf("succeed to post {uid: %d, mid: %d}", uid, mid)
	}
	return err
}

func (d *Ingester) retry() {
	for {
		select {
		case <-d.done:
			return
		case p := <-d.retryCh:
			interval := time.Since(p.last).Seconds()
			if interval < 60 {
				time.Sleep(time.Duration((60 - interval)) * time.Second)
			}
			err := postLongText(p.uid, p.mid)
			if err != nil && p.retries > 0 {
				p.retries--
				p.last = time.Now()
				select {
				case d.retryCh <- p:
				default:
					log.Printf("retry channel if full so {uid: %d, mid: %d} is dropped", p.uid, p.mid)
				}
			} else if p.retries == 0 {
				log.Printf("reach max retry, this message is dropped: {uid: %d, mid: %d}", p.uid, p.mid)
			}
		}
	}
}

func (d *Ingester) drain(done chan struct{}) {
	for p := range d.retryCh {
		err := postLongText(p.uid, p.mid)
		if err != nil {
			log.Printf("failed to drain message: {uid: %d, mid: %d}", p.uid, p.mid)
		}
	}

	close(done)
}
