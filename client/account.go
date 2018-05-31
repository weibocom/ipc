package client

import (
	"errors"

	"github.com/weibocom/ipc/config"
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
	//return c.steem.Condenser.GetAccountCount()
	count, err := c.store.GetAccountCount()
	return uint32(count), err
}

func (c *client) GetAccountPostCount(name string) (int, error) {
	accs, err := c.steem.Condenser.GetAccounts([]string{name})
	if err != nil || len(accs) == 0 {
		return 0, err
	}

	return int(accs[0].PostCount.Int64()), nil
}

func (c *client) LookupAccount(name string) (*model.Account, error) {
	// exist both local db and chain

	// check whether this account exists in chain
	accounts, err := c.steem.Condenser.LookupAccountNames([]string{name})
	if err != nil {
		return nil, err
	}
	if len(accounts) == 0 {
		return nil, errors.New("account not found in chain")
	}

	// check store
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
	publicKey := wif.PublicKey().String()

	creator := config.GetCreator()
	fee := config.GetCreateAccountFee()
	err = c.steem.CreateAccount(creator, name, fee, publicKey, publicKey, publicKey, publicKey, meta)
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
