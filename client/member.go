package client

import (
	"fmt"

	"github.com/weibocom/ipc/keys"
	"github.com/weibocom/ipc/model"

	"github.com/weibocom/ipc/config"
	"github.com/weibocom/ipc/store"
)

// Members get all members in chain that contains all accounts.
func (c *client) Members() ([]*model.Member, error) {
	witnesses, err := c.steem.Condenser.GetWitnesses([]uint32{0})
	if err != nil {
		return nil, err
	}

	var members = make([]*model.Member, len(witnesses))
	for i, w := range witnesses {
		m := &model.Member{}
		m.ID = w.ID.Int64()
		m.Name = w.Owner
		m.SigningKey = w.SigningKey
		m.CreatedAt = *w.Created.Time
		members[i] = m
	}
	return members, nil
}

// AddMember add this member as witness.
func (c *client) AddMember(name string) (*model.Member, error) {
	ok, err := c.store.ExistMember(name)
	if err != nil {
		return nil, err
	}
	if ok {
		return nil, fmt.Errorf("witness %s is already existed", name)
	}

	// check the account exists
	acc, err := c.lookupAccount(name)
	if err != nil && err != store.ErrNonExist {
		return nil, err
	}

	// create this account if it doesn't exist
	if err == store.ErrNonExist {
		acc, err = c.CreateAccount(name, `{"meta":"`+name+`"}`)
		if err != nil {
			return nil, err
		}
	}
	privateKeys := [][]byte{acc.WIF.PrivateKey().Serialize()}
	//privateKeys = keys.GetPrivateKeys()
	wif := config.GetWIFs()[0]
	dpk, _ := keys.DecodeWIF(wif)
	pk := dpk.PrivateKey().Public().String()
	fmt.Println("initminer private_key: ", wif)
	fmt.Println("initminer public_key: ", pk)
	fmt.Println("account wif: ", acc.WIF.String())
	fmt.Println("account block_sign_key: ", acc.WIF.PrivateKey().Public().String())

	err = c.steem.AddWitness(privateKeys, name, acc.WIF.PrivateKey().Public().String(), config.GetURL(), config.GetCreateAccountFee())
	if err != nil {
		return nil, err
	}

	w, err := c.steem.Condenser.GetWitnessesByAccount(name)
	if err != nil {
		return nil, err
	}

	if w == nil {
		return nil, fmt.Errorf("sent update_witness for %s but can't lookup", name)
	}

	m := &model.Member{Name: name}
	m.ID = w.ID.Int64()
	m.SigningKey = w.SigningKey
	m.CreatedAt = *w.Created.Time

	//save to store
	return m, c.store.SaveMember(m)
}
