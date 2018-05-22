package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/weibocom/ipc/keys"
	ipcmodel "github.com/weibocom/ipc/model"
	"github.com/weibocom/ipc/web2/model"
)

var (
	ErrCompanyIsNotInConsortium = errors.New("not a memeber of consortium")
	ErrUserNotExist             = errors.New("user does not exist")
	ErrUserAlreadyExist         = errors.New("user has already existed")
)

func generateUniqueAccount(company string, user int64) string {
	return fmt.Sprintf("%s-%d", company, user)
}

func splitCompanyAccount(author string) (company string, uidStr string) {
	// must
	fields := strings.SplitN(author, "-", 2)

	if len(fields) == 2 {
		company = fields[0]
		uidStr = fields[1]
	} else {
		uidStr = author
	}

	return
}

func getUserID(name string) int64 {
	i := strings.LastIndex(name, "-")
	if i == -1 {
		return 0
	}

	id, _ := strconv.ParseInt(name[i+1:], 10, 64)
	return id
}

func RegisterAccount(company string, user int64) (publicKey string, privateKey string, err error) {
	acc, err := ipcClient.CreateAccount(generateUniqueAccount(company, user), `{"company":"`+company+`"}`)
	if err != nil {
		return "", "", err
	}

	wif := acc.WIF
	userWIF, err := keys.DecodeWIF(wif)
	if err != nil {
		return "", "", err
	}

	return userWIF.PrivateKey().HexString(), userWIF.PrivateKey().Public().String(), nil
}

func GetUser(company string, uid int64) (*model.User, error) {
	author := generateUniqueAccount(company, uid)

	acc, err := ipcClient.LookupAccount(author)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:        uid,
		Company:   company,
		CreatedAt: acc.CreatedAt,
	}
	wif := acc.WIF
	userWIF, err := keys.DecodeWIF(wif)
	if err == nil {
		user.PrivateKey = userWIF.PrivateKey().HexString()
		user.PublicKey = userWIF.PrivateKey().Public().String()
	}

	user.PostCount, _ = ipcClient.GetAccountPostCount(author)
	return user, nil
}

func GetUsers(company string, page int, pagesize int, uid int64) ([]*model.User, error) {
	offset := (page - 1) * pagesize

	var accounts []*ipcmodel.Account
	var err error

	if uid == -1 {
		accounts, err = ipcClient.GetAccounts(company, offset, pagesize)
		if err != nil {
			return nil, err
		}
	} else { // 查询单个账号
		author := generateUniqueAccount(company, uid)
		acc, err := ipcClient.LookupAccount(author)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, acc)
	}

	var users = make([]*model.User, 0, len(accounts))
	for _, acc := range accounts {
		user := &model.User{
			ID:        getUserID(acc.Name),
			Company:   company,
			CreatedAt: acc.CreatedAt,
		}

		wif := acc.WIF
		userWIF, err := keys.DecodeWIF(wif)
		if err == nil {
			user.PrivateKey = userWIF.PrivateKey().HexString()
			user.PublicKey = userWIF.PrivateKey().Public().String()
		}

		// user.PostCount = ...

		users = append(users, user)
	}

	return users, nil
}

func UserCount() (int64, error) {
	count, err := ipcClient.AccountCount()
	return int64(count), err
}
