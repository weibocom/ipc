package client

import (
	"github.com/weibocom/ipc/config"
	"github.com/weibocom/ipc/keys"
)

// TODO
func (c *client) checkAccount(name string) error {
	return nil
}

// TODO
func (c *client) saveAccount(a *Account) error {
	return nil
}

func (c *client) lookupAccount(name string) (*Account, error) {
	return nil, nil
}

func (c *client) CreateAccount(name string, meta string) (*Account, error) {
	if err := c.checkAccount(name); err != nil {
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
	c.saveAccount(account)
	return account, nil
}
