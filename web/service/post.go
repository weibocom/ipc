package service

import (
	"strconv"

	"github.com/weibocom/ipc/model"
	"github.com/weibocom/ipc/store"
	webmodel "github.com/weibocom/ipc/web/model"
)

func SendBlog(userID int64, msgID int64, company string, title string, content string, currentTs int64) (model.DNA, error) {
	if company == "" {
		company = "weibo"
	}
	companyAbbr := Companies[company]

	if companyAbbr == "" {
		return nil, ErrCompanyIsNotInConsortium
	}

	author := generateSteemAccount(Companies[company], userID)

	_, err := IPCClient.LookupAccount(author)
	if err == store.ErrNonExist {
		_, err = IPCClient.CreateAccount(author, `{"meta":"`+company+`"}`)
		if err != nil {
			return nil, err
		}
	}

	return IPCClient.Post(author, title, []byte(content), strconv.FormatInt(msgID, 10), []string{"test"})
}

func GetContent(company string, userID int64, dna string) (*model.Post, error) {
	companyAbbr := Companies[company]

	if companyAbbr == "" {
		return nil, ErrCompanyIsNotInConsortium
	}

	author := generateSteemAccount(Companies[company], userID)

	return IPCClient.LookupPost(author, model.DNA(dna))
}

func GetContentByMsgID(company string, userID int64, id int64) (*model.Post, error) {
	companyAbbr := Companies[company]

	if companyAbbr == "" {
		return nil, ErrCompanyIsNotInConsortium
	}

	author := generateSteemAccount(Companies[company], userID)

	// Trick, demo把weibo长文的msgid放在了uri字段中
	return IPCClient.LookupPostByURI(author, strconv.FormatInt(id, 10))
}

func GetBlogEntrys(company string, userID int64, dna model.DNA, limit uint16) ([]*webmodel.BlogEntry, error) {
	companyAbbr := Companies[company]
	if companyAbbr == "" {
		return nil, ErrCompanyIsNotInConsortium
	}
	author := generateSteemAccount(Companies[company], userID)

	entries, err := IPCClient.GetPosts(author, dna, 20)
	if err != nil {
		return nil, err
	}

	var blogs []*webmodel.BlogEntry
	for _, e := range entries {
		msgid, _ := strconv.ParseInt(e.URI, 10, 64)
		entry := &webmodel.BlogEntry{
			Author:   e.Author,
			Permlink: e.DNA,
			EntryID:  uint32(msgid),
		}

		blogs = append(blogs, entry)
	}

	return blogs, nil
}

func GetLatestMid() (*model.Post, error) {
	return IPCClient.GetLatestPost()
}
