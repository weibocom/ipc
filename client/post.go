package client

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/weibocom/ipc/content"
	"github.com/weibocom/ipc/model"
	"github.com/weibocom/ipc/signature"
	"github.com/weibocom/ipc/util"
)

// TODO
// snapshot 1. 加密存储； 2. 返回存储后的唯一id。通常是snapshot的digest
func (c *client) snapshot(account *model.Account, author string, title string, content []byte, uri string) ([]byte, model.DNA, error) {
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

	post := &model.Post{
		Author:  author,
		Title:   title,
		Content: content, // TODO: 加密
		URI:     uri,
		Digest:  digest,
		DNA:     dna,
	}

	err = c.store.SavePost(post)
	return digest, dna, err
}

func (c *client) sign(a *model.Account, digest []byte) (model.DNA, error) {
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
	return model.DNA(hex.EncodeToString(sigs[0])), nil
}

func (c *client) Post(author string, title string, content []byte, uri string, tags []string) (model.DNA, error) {
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

func (c *client) LookupContent(dna model.DNA) (model.Content, error) {
	post, err := c.store.LoadPost(dna)
	if err != nil {
		return nil, err
	}

	return post.Content, nil
}

func (c *client) Verify(author string, dna model.DNA) (bool, error) {
	post, err := c.store.LoadPost(dna)
	if err != nil {
		return false, err
	}

	return author == post.Author, nil
}

func (c *client) CheckSimilar(a, b model.DNA) (float64, error) {
	post1, err := c.store.LoadPost(a)
	if err != nil {
		return 0, err
	}
	post2, err := c.store.LoadPost(b)
	if err != nil {
		return 0, err
	}

	return content.Similarity(string(post1.Content), string(post2.Content)), nil
}
