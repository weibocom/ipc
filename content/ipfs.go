package content

import (
	"io/ioutil"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

type IPFSClient struct {
	url    string
	client *shell.Shell
}

func NewIPFSClient(u string) *IPFSClient {
	s := &IPFSClient{
		url: u,
	}

	client := shell.NewShell(u)
	s.client = client

	return s
}

func (c *IPFSClient) AddContent(content string) (string, error) {
	return c.client.Add(strings.NewReader(content))
}

func (c *IPFSClient) GetContent(hash string) (string, error) {
	r, err := c.client.Cat(hash)
	if err != nil {
		return "", err
	}
	defer r.Close()

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}
