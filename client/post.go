package client

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/weibocom/ipc/signature"
	"github.com/weibocom/ipc/util"
)

type Post struct {
	Author  string
	Title   string
	Content []byte
	URI     string
	Digest  []byte
	DNA     DNA
}

// TODO
// snapshot 1. 加密存储； 2. 返回存储后的唯一id。通常是snapshot的digest
func (c *client) snapshot(account *Account, author string, title string, content []byte, uri string) ([]byte, DNA, error) {
	sha := sha256.New()
	sha.Write(util.String2Bytes(author))
	sha.Write(util.String2Bytes(title))
	sha.Write(content)
	sha.Write(util.String2Bytes(uri))
	digest := sha.Sum(nil)

	dna, err := c.sign(account, digest)

	if err != nil {
		return digest, nil, err
	}

	post := &Post{
		Author:  author,
		Title:   title,
		Content: content, // TODO: 加密
		URI:     uri,
		Digest:  digest,
		DNA:     dna,
	}

	data, err := util.ToJSON(post)
	if err != nil {
		return digest, dna, err
	}

	err = c.store.Save(PostStoreType, string(dna), data)
	return digest, dna, err
}

func (c *client) sign(a *Account, digest []byte) (DNA, error) {
	// TODO: reuse signature.Sign but it is for multiple keys
	privKeys := [][]byte{a.WIF.PrivateKey().Serialize()}
	s := signature.NewSignature()
	sigs, err := s.Sign(privKeys, digest)
	if err != nil {
		return nil, err
	}

	if len(sigs) != 1 {
		return nil, errors.New("must be one sigurature")
	}
	return DNA(hex.EncodeToString(sigs[0])), nil
}

func (c *client) Post(author string, title string, content []byte, uri string, tags []string) (DNA, error) {
	account, err := c.lookupAccount(author)
	if err != nil {
		return nil, err
	}

	digest, dna, err := c.snapshot(account, author, title, content, uri)
	if err != nil {
		return nil, err
	}

	body := hex.EncodeToString(digest)

	privateKeys := [][]byte{account.WIF.PrivateKey().Serialize()}
	c.steem.Post(privateKeys, author, title, body, dna.ID(), dna.ID(), "", tags)

	return dna, nil
}
