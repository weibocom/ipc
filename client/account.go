package client

import (
	"errors"

	"github.com/weibocom/ipc/config"
	"github.com/weibocom/ipc/keys"
	"github.com/weibocom/ipc/util"
)

// TODO
func (c *client) checkAccount(name string) (bool, error) {
	return c.store.Exist(AccountStoreType, name)
}

func (c *client) saveAccount(a *Account) error {
	v, err := util.ToJSON(a)
	if err != nil {
		return err
	}
	return c.store.Save(AccountStoreType, a.Name, v)
}

func (c *client) lookupAccount(name string) (*Account, error) {
	v, err := c.store.Load(AccountStoreType, name)
	if err != nil {
		return nil, err
	}
	a := &Account{}
	err = util.FromJSON(v, a)
	return a, err
}

func (c *client) CreateAccount(name string, meta string) (*Account, error) {
	exist, err := c.checkAccount(name)
	if exist {
		return nil, errors.New("account is already existed")
	}
	if err != nil {
		return nil, err
	}

	wif, err := keys.GenerateWIF()
	if err != nil {
		return nil, err
	}
	publicKey := wif.PublicKey().String()

	creator := config.GetCreator()
	fee := config.GetCreateAccountFee()
	err = c.steem.CreateAccount(creator, name, fee, publicKey, publicKey, publicKey, publicKey, meta)
	if err != nil {
		return nil, err
	}

	account := &Account{
		Name: name,
		WIF:  wif,
	}
	err = c.saveAccount(account)
	return account, err
}
