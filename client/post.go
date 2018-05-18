package client

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/weibocom/ipc/content"
	"github.com/weibocom/ipc/keys"
	"github.com/weibocom/ipc/model"
	"github.com/weibocom/ipc/signature"
	"github.com/weibocom/ipc/util"
)

// TODO
// snapshot 1. 加密存储； 2. 返回存储后的唯一id。通常是snapshot的digest
func (c *client) snapshot(account *model.Account, mid int64, author string, title string, content []byte, uri string) ([]byte, model.DNA, error) {
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
		MSGID:     mid,
		Author:    author,
		Title:     title,
		Content:   string(content), // TODO: 加密
		URI:       uri,
		Digest:    hex.EncodeToString(digest),
		DNA:       dna.ID(),
		CreatedAt: time.Now(),
	}

	err = c.store.SavePost(post)
	return digest, dna, err
}

func (c *client) sign(a *model.Account, digest []byte) (model.DNA, error) {
	accWif, err := keys.DecodeWIF(a.WIF)
	if err != nil {
		return nil, err
	}

	privKeys := [][]byte{accWif.PrivateKey().Serialize()}
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

func (c *client) Post(author string, mid int64, title string, content []byte, uri string, tags []string) (model.DNA, error) {
	account, err := c.lookupAccount(author)
	if err != nil {
		return nil, err
	}

	digest, dna, err := c.snapshot(account, mid, author, title, content, uri)
	if err != nil {
		return nil, err
	}

	body := hex.EncodeToString(digest)

	accWif, err := keys.DecodeWIF(account.WIF)
	if err != nil {
		return nil, err
	}
	privateKeys := [][]byte{accWif.PrivateKey().Serialize()}
	_, err = c.steem.Post(privateKeys, author, title, body, dna.ID(), dna.ID(), "", tags)

	return dna, err
}

func (c *client) PostAsync(author string, mid int64, title string, content []byte, uri string, tags []string) (model.DNA, error) {
	account, err := c.lookupAccount(author)
	if err != nil {
		return nil, err
	}

	digest, dna, err := c.snapshot(account, mid, author, title, content, uri)
	if err != nil {
		return nil, err
	}

	body := hex.EncodeToString(digest)

	accWif, err := keys.DecodeWIF(account.WIF)
	if err != nil {
		return nil, err
	}
	privateKeys := [][]byte{accWif.PrivateKey().Serialize()}
	err = c.steem.PostAsync(privateKeys, author, title, body, dna.ID(), dna.ID(), "", tags)

	return dna, err
}

func (c *client) LookupContent(dna model.DNA) (model.Content, error) {
	post, err := c.store.LoadPost(dna)
	if err != nil {
		return nil, err
	}

	return []byte(post.Content), nil
}

func (c *client) LookupPost(auther string, dna model.DNA) (*model.Post, error) {
	content, err := c.steem.Condenser.GetContent(auther, dna.ID())
	if err != nil {
		return nil, err
	}

	cc, err := c.LookupContent(dna)

	return &model.Post{
		DNA:     dna.ID(),
		Author:  auther,
		Title:   content.Title,
		Content: string(cc), // TODO: 从IPFS获取内容
		URI:     content.URL,
		Digest:  content.Body,
	}, err

}

func (c *client) GetLatestPost() (*model.Post, error) {
	return c.store.GetLatestPost()
}

func (c *client) LookupPostByMsgID(author string, mid int64) (*model.Post, error) {
	return c.store.GetPostByMsgID(author, mid)
}

func (c *client) LookupPostByDNA(dna model.DNA) (*model.Post, error) {
	return c.store.GetPostByDNA(dna)
}

func (c *client) Verify(author string, dna model.DNA) (bool, error) {
	// query this post in chain
	content, err := c.steem.Condenser.GetContent(author, dna.ID())
	if err != nil {
		return false, err
	}

	return content != nil, nil
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

	return content.Similarity(post1.Content, post2.Content), nil
}
