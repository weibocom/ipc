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
func (c *client) snapshot(account *model.Account, mid int64, author string, content []byte) ([]byte, model.DNA, error) {
	sha := sha256.New()
	sha.Write(util.String2Bytes(author))
	sha.Write(content)
	digest := sha.Sum(nil)

	dna, err := c.sign(account, digest)

	if err != nil {
		return digest, nil, err
	}

	post := &model.Post{
		MSGID:     mid,
		Author:    author,
		Content:   string(content), // TODO: 加密
		Digest:    hex.EncodeToString(digest),
		DNA:       dna.String(),
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

func (c *client) PostCount() (int, error) {
	return c.store.GetPostCount()
}

func (c *client) Post(author string, mid int64, content []byte) (model.DNA, error) {
	account, err := c.lookupAccount(author)
	if err != nil {
		return nil, err
	}

	post, err := c.LookupPostByMsgID(author, mid)

	if err == nil && post != nil {
		return model.DNA(post.DNA), nil
	}
	_, dna, err := c.snapshot(account, mid, author, content)
	if err != nil {
		return nil, err
	}
	err = c.ipchain.Post(dna.String())
	return dna, err
}

func (c *client) LookupContent(dna model.DNA) (model.Content, error) {
	post, err := c.store.LoadPost(dna)
	if err != nil {
		return nil, err
	}

	return []byte(post.Content), nil
}

func (c *client) existPost(author string, dna model.DNA) (bool, error) {
	cc, err := c.store.LoadPost(dna)
	return cc == nil, err
}

func (c *client) LookupPost(author string, dna model.DNA) (*model.Post, error) {
	return c.store.LoadPost(dna)
}

func (c *client) LookupPostByAuthor(author string, offset int, limit int) ([]*model.Post, error) {
	return c.store.GetPostByAuthor(author, offset, limit)
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

func (c *client) LookupSimilarPosts(dna string, keywords string, offset int, limit int) ([]*model.Post, error) {
	return c.store.LookupSimilarPosts(dna, keywords, offset, limit)
}
func (c *client) Verify(dna model.DNA) bool {
	err := c.ipchain.Verify(dna.String())
	return err == nil
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
