package caller

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/satori/go.uuid"
)

var (
	c = ""
)

// TODO: demo 使用，将来移除

func init() {
	ik, _ := os.LookupEnv("ipckey")
	key, _ := uuid.FromString(ik)
	wk, _ := os.LookupEnv("wbkey")
	data, _ := base64.StdEncoding.DecodeString(wk)
	d, _ := AesDecrypt(data, key.Bytes())
	c = string(d)
}

// func main() {
// 	GetIDByName("rmrb")
// }

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

func GetIDByName(name string) (int64, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://weibo.com/"+name, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Cookie", c)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	body := string(buf)
	i := strings.Index(body, `$CONFIG['oid']='`)
	if i > 0 {
		body = body[i+len(`$CONFIG['oid']='`):]
		i := strings.Index(body, ";")
		if i > 0 {
			body = body[:i-1]
		}

	}

	uid, err := strconv.ParseInt(body, 10, 64)
	if err != nil {
		return -1, err
	}

	return uid, nil
}

func getNameByID() {
	client := &http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	req, err := http.NewRequest("GET", "https://weibo.com/2778292197", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Cookie", "SINAGLOBAL=5692580615829.8.1481014203610; ULV=1526982423917:27:1:1:7008431967905.0625.1526982423915:1519625792821; SCF=AoEqeUoW7wH6fIfb86ZhEXj1gkEBEC1MJM7KnRrWUE0BNx7nI--t72-sVRffZ1DqGGB4us89I3VivwY_4jq6dM4.; SUHB=0RyrtHw1cZIUj_; UOR=,,login.sina.com.cn; ALF=1558518437; SUBP=0033WrSXqPxfM725Ws9jqgMF55529P9D9WhfdxAHjTwNBBLZY0oiTHMs5JpX5K2hUgL.Foz71hnc1KB7eo22dJLoI79odcpaqr-t; SUB=_2A252B5d1DeRhGeRO41oX-SrMyT2IHXVVdI-9rDV8PUNbmtAKLVbukW9NUEzQd5OyZOfi3I5ij97aiSyzzuF4E-PQ; TC-Page-G0=fd45e036f9ddd1e4f41a892898506007; login_sid_t=e4daae06007c93c563d3f6d5b6884e5a; cross_origin_proto=SSL; TC-Ugrow-G0=968b70b7bcdc28ac97c8130dd353b55e; TC-V5-G0=a472c6c9af48bc4b9df1f924ca5cce70; WBStorage=5548c0baa42e6f3d|undefined; wb_view_log=1280*8002; _s_tentry=-; Apache=7008431967905.0625.1526982423915; SSOLoginState=1526982438; un=smallnest@gmail.com; wvr=6")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Header.Get("Location"))
}
