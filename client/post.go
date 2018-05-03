package client

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/rfguri/bowsim"
	"github.com/weibocom/ipc/signature"
	"github.com/weibocom/ipc/util"
	"github.com/yanyiwu/gojieba"
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

	err = c.store.Save(PostStoreType, dna.ID(), data)
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

func (c *client) LookupContent(dna DNA) (Content, error) {
	return c.store.Load(PostStoreType, dna.ID())
}

func (c *client) Verify(author string, dna DNA) (bool, error) {
	v, err := c.store.Load(PostStoreType, dna.ID())
	if err != nil {
		return false, err
	}

	post := &Post{}
	err = util.FromJSON(v, post)
	if err != nil {
		return false, err
	}

	return author == post.Author, nil
}

func (c *client) CheckSimilar(a, b DNA) (float64, error) {
	v1, err := c.store.Load(PostStoreType, a.ID())
	if err != nil {
		return 0, err
	}
	post1 := &Post{}
	err = util.FromJSON(v1, post1)
	if err != nil {
		return 0, err
	}

	v2, err := c.store.Load(PostStoreType, b.ID())
	if err != nil {
		return 0, err
	}
	post2 := &Post{}
	err = util.FromJSON(v2, post2)
	if err != nil {
		return 0, err
	}

	return similarity(string(post1.Content), string(post2.Content)), nil
}

func similarity(a, b string) float64 {
	x := gojieba.NewJieba()
	defer x.Free()

	words1 := x.Cut(a, true)
	words2 := x.Cut(b, true)

	return bowsim.Get(strings.Join(words1, " "), strings.Join(words2, " "))
}
