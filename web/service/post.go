package service

import (
	"fmt"
	"strconv"

	"github.com/weibocom/ipc/content"
	"github.com/weibocom/ipc/model"
	"github.com/weibocom/ipc/store"
	webmodel "github.com/weibocom/ipc/web/model"
)

func AddPost(company string, uid int64, mid int64, title string, content string, currentTs int64) (model.DNA, error) {
	author := generateUniqueAccount(company, uid)
	_, err := ipcClient.LookupAccount(author)
	if err == store.ErrNonExist {
		_, err = ipcClient.CreateAccount(author, `{"company":"`+company+`"}`)
		if err != nil {
			return nil, err
		}
	}

	return ipcClient.Post(author, mid, title, []byte(content), strconv.FormatInt(mid, 10), []string{"test"})
}

func GetContentByMsgID(company string, uid int64, mid int64) (*model.Post, error) {
	author := generateUniqueAccount(company, uid)
	post, err := ipcClient.LookupPostByMsgID(author, mid)

	if post != nil {
		_, post.Author = splitCompanyAccount(post.Author)
	}

	return post, err
}

func GetContentByDNA(dna string) (*model.Post, error) {
	post, err := ipcClient.LookupPostByDNA(model.DNA(dna))

	if post != nil {
		_, post.Author = splitCompanyAccount(post.Author)
	}

	return post, err
}

func GetLatestPost() (*model.Post, error) {
	post, err := ipcClient.GetLatestPost()
	if post != nil {
		_, post.Author = splitCompanyAccount(post.Author)
	}

	return post, err
}

func GetUserPosts(company string, uid int64, page int, pagesize int) (posts []*model.Post, postCount int, err error) {
	offset := (page - 1) * pagesize

	author := generateUniqueAccount(company, uid)

	postCount, err = ipcClient.GetAccountPostCount(author)
	if err != nil {
		return
	}

	posts, err = ipcClient.LookupPostByAuther(author, offset, pagesize)

	for _, p := range posts {
		_, p.Author = splitCompanyAccount(p.Author)
	}
	return posts, postCount, err
}

func GetSimilarPostsByDNA(dna string, c string, keywords string, page int, pagesize int) ([]*webmodel.Post, error) {
	offset := (page - 1) * pagesize
	posts, err := ipcClient.LookupSimilarPosts(dna, keywords, offset, pagesize)

	var webposts []*webmodel.Post

	for _, p := range posts {
		_, p.Author = splitCompanyAccount(p.Author)
		pp := &webmodel.Post{}
		pp.Post = p

		pp.Similarity = fmt.Sprintf("%.2f", content.Similarity(c, p.Content)*100)
		webposts = append(webposts, pp)
	}
	return webposts, err
}

func PostCount() (int64, error) {
	count, err := ipcClient.PostCount()
	return int64(count), err
}
