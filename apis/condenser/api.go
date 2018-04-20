package condenser

import (
	"encoding/json"

	"github.com/weibocom/steem-rpc/interfaces"
	"github.com/weibocom/steem-rpc/internal/call"
)

const (
	APIID = "condenser_api"
)

var emptyParams = []string{}

type API struct {
	caller interfaces.Caller
}

func NewAPI(caller interfaces.Caller) *API {
	return &API{caller}
}

/*
   // Witnesses
   (get_witnesses)
   (get_witness_by_account)
   (get_witnesses_by_vote)
   (lookup_witness_accounts)
   (get_witness_count)
   (get_active_witnesses)
   (get_miner_queue)
*/

// GetWitnesses implements get_witnesses
func (api *API) GetWitnesses(id []uint32) ([]*Witness, error) {
	raw, err := call.Raw(api.caller, APIID+".get_witnesses", [][]uint32{id})
	if err != nil {
		return nil, err
	}
	var resp []*Witness
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *API) GetWitnessesByAccount(author string) (*Witness, error) {
	raw, err := call.Raw(api.caller, APIID+".get_witness_by_account", []string{author})
	if err != nil {
		return nil, err
	}
	var resp *Witness
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *API) LookupWitnessAccounts(author string, limit uint) ([]string, error) {
	raw, err := call.Raw(api.caller, APIID+".lookup_witness_accounts", []string{author})
	if err != nil {
		return nil, err
	}
	var resp []string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *API) GetWitnessCount() (uint32, error) {
	raw, err := call.Raw(api.caller, APIID+".get_witness_count", emptyParams)
	if err != nil {
		return 0, err
	}
	var resp uint32
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return 0, err
	}
	return resp, nil
}

func (api *API) GetActiveWitnesses() ([]string, error) {
	raw, err := call.Raw(api.caller, APIID+".get_active_witnesses", emptyParams)
	if err != nil {
		return nil, err
	}
	var resp []string
	if err := json.Unmarshal([]byte(*raw), &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
