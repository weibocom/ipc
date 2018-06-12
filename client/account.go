package client

import (
	"github.com/weibocom/ipc/keys"
	"github.com/weibocom/ipc/model"
)

func (c *client) checkAccount(name string) (bool, error) {
	return c.store.ExistAccount(name)
}

func (c *client) saveAccount(a *model.Account) error {
	return c.store.SaveAccount(a)
}

func (c *client) lookupAccount(name string) (*model.Account, error) {
	return c.store.LoadAccount(name)
}

func (c *client) GetAccounts(company string, offset int, limit int) ([]*model.Account, error) {
	return c.store.GetAccounts(company, offset, limit)
}

func (c *client) AccountCount() (uint32, error) {
	count, err := c.store.GetAccountCount()
	return uint32(count), err
}

func (c *client) GetAccountPostCount(name string) (int, error) {
	return c.store.GetAccountCount()
}

func (c *client) LookupAccount(name string) (*model.Account, error) {
	return c.store.LoadAccount(name)
}

func (c *client) CreateAccount(name string, meta string) (*model.Account, error) {
	exist, err := c.checkAccount(name)
	if exist {
		return nil, ErrAccountAlreadyExist
	}
	if err != nil {
		return nil, err
	}

	wif, err := keys.GenerateWIF()
	if err != nil {
		return nil, err
	}
	account := &model.Account{
		Name: name,
		WIF:  wif.String(),
	}
	err = c.saveAccount(account)
	return account, err
}
