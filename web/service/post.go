package service

import (
	"strconv"

	"github.com/weibocom/ipc/model"
	"github.com/weibocom/ipc/store"
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

	return IPCClient.Post(author, msgID, title, []byte(content), strconv.FormatInt(msgID, 10), []string{"test"})
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

	return IPCClient.LookupPostByMsgID(author, id)
}

func GetLatestMid() (*model.Post, error) {
	return IPCClient.GetLatestPost()
}
