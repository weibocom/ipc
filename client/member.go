package client

import (
	"fmt"

	"github.com/weibocom/ipc/config"
	"github.com/weibocom/ipc/store"
	"github.com/weibocom/ipc/util"
)

// Members get all members in chain that contains all accounts.
func (c *client) Members() ([]*Member, error) {
	witnesses, err := c.steem.Condenser.GetWitnesses([]uint32{0})
	if err != nil {
		return nil, err
	}

	var members = make([]*Member, len(witnesses))
	for i, w := range witnesses {
		m := &Member{}
		m.ID = w.ID.Int64()
		m.Name = w.Owner
		m.SigningKey = w.SigningKey
		m.CreatedAt = *w.Created.Time
		members[i] = m
	}
	return members, nil
}

// AddMember add this member as witness.
func (c *client) AddMember(name string) (*Member, error) {
	ok, err := c.store.Exist(MemberStoreType, name)
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
	err = c.steem.AddWitness(privateKeys, name, acc.WIF.PublicKey().String(), config.GetURL(), config.GetCreateAccountFee())
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

	m := &Member{Name: name}
	m.ID = w.ID.Int64()
	m.SigningKey = w.SigningKey
	m.CreatedAt = *w.Created.Time

	//save to store
	v, err := util.ToJSON(m)
	if err != nil {
		return nil, err
	}
	return m, c.store.Save(MemberStoreType, name, v)
}
