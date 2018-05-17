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

func generateUniqueAccount(company string, user int64) string {
	return fmt.Sprintf("%s-%d", company, user)
}

func RegisterAccount(company string, user int64) error {
	_, err := IPCClient.CreateAccount(generateUniqueAccount(company, user), `{"company":"`+company+`"}`)
	return err
}

func UserCount() (int64, error) {
	count, err := IPCClient.AccountCount()
	return int64(count), err
}
