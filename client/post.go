package client

import "encoding/hex"

// TODO
// snapshot 1. 加密存储； 2. 返回存储后的唯一id。通常是snapshot的digest
func (c *client) snapshot(author string, title string, content []byte, uri string) ([]byte, error) {
	return []byte{}, nil
}

func (c *client) sign(a *Account, digest []byte) (*DNA, error) {
	return nil, nil
}

func (c *client) Post(author string, title string, content []byte, uri string, tags []string) (*DNA, error) {
	account, err := c.lookupAccount(author)
	if err != nil {
		return nil, err
	}

	digest, err := c.snapshot(author, title, content, uri)

	if err != nil {
		return nil, err
	}

	dna, err := c.sign(account, digest)

	if err != nil {
		return nil, err
	}

	body := hex.EncodeToString(digest)

	c.steem.Post(author, title, body, dna.ID(), "", "", tags)

	return nil, nil
}
