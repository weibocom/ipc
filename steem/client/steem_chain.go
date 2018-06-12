package client

import (
	"errors"
	"time"

	"github.com/weibocom/ipc/interfaces"
)

type AsyncPostCall struct {
	DNA         string
	Error       error
	MaxWaitTime time.Duration

	done chan struct{}
}

type Steem struct {
	steem      *Client
	submitter  string
	privateKey []byte
	company    string

	asyncPostCallChan chan *AsyncPostCall
	done              chan struct{}
}

func NewSteemClient(cc interfaces.CallCloser, submitter string, privateKey []byte, company string) *Steem {
	steem, err := NewClient(cc)
	if err != nil {
		panic(err)
	}
	s := &Steem{
		steem:             steem,
		submitter:         submitter,
		privateKey:        privateKey,
		company:           company,
		asyncPostCallChan: make(chan *AsyncPostCall, 10000),
		done:              make(chan struct{}),
	}

	go s.handleAsyncPostCall()

	return s
}

func (s *Steem) handleAsyncPostCall() {
	for {
		select {
		case <-s.done:
			return
		case call := <-s.asyncPostCallChan:

			start := time.Now()
			var err error
			for {
				err = s.Verify(call.DNA)
				if err != nil {
					time.Sleep(200 * time.Millisecond)
					if time.Since(start) > call.MaxWaitTime {
						break
					}
				} else {
					break
				}
			}
			call.Error = err
			close(call.done)
		}
	}
}

func (s *Steem) Post(dna string) error {
	err := s.postAsync(dna)
	if err != nil {
		return err
	}

	call := &AsyncPostCall{
		DNA:         dna,
		MaxWaitTime: 30 * time.Second,
		done:        make(chan struct{}),
	}

	s.asyncPostCallChan <- call
	<-call.done

	return call.Error
}

func (s *Steem) postAsync(dna string) error {
	return s.steem.PostAsync([][]byte{s.privateKey}, s.submitter,
		"title", dna, dna, s.company, "", []string{`{}`})
}

func (s *Steem) Verify(dna string) error {
	content, err := s.steem.Condenser.GetContent(s.submitter, dna)
	if err != nil {
		return err
	}

	if content == nil {
		return errors.New("content not found")
	}
	return nil
}

func (s *Steem) Close() error {
	close(s.done)
	return s.steem.Close()
}
