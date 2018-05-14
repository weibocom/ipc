package service

import (
	"errors"
	"fmt"
)

var (
	ErrCompanyIsNotInConsortium = errors.New("not a memeber of consortium")
	ErrUserNotExist             = errors.New("user does not exist")
	ErrUserAlreadyExist         = errors.New("user has already existed")
)
var (
	Companies = map[string]string{
		"weibo":    "wb",
		"zhihu":    "zh",
		"facebook": "fb",
	}
)

func generateSteemAccount(company string, user int64) string {
	// %d-%s is invalid
	return fmt.Sprintf("%s-%d", company, user)
}

func RegisterIPAccount(company string, user int64) error {
	if company == "" {
		company = "weibo"
	}
	companyAbbr := Companies[company]

	if companyAbbr == "" {
		return ErrCompanyIsNotInConsortium
	}

	_, err := IPCClient.CreateAccount(generateSteemAccount(companyAbbr, user), `{"meta":"gosteem"}`)
	return err
}

func GetUserCount() (int64, error) {
	count, err := IPCClient.AccountCount()
	return int64(count), err
}
