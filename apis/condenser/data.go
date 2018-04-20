package condenser

import "github.com/weibocom/steem-rpc/types"

type ChainProperties struct {
	AccountCreationFee string     `json:"account_creation_fee"`
	MaximumBlockSize   *types.Int `json:"maximum_block_size"`
	SbdInterestRate    *types.Int `json:"sbd_interest_rate"`
}

type CurrentMedianHistoryPrice struct {
	Base  string `json:"base"`
	Quote string `json:"quote"`
}

type Witness struct {
	ID                    *types.Int                 `json:"id"`
	Owner                 string                     `json:"owner"`
	Created               *types.TimePointSeconds    `json:"created"`
	URL                   string                     `json:"url"`
	Votes                 *types.Int                 `json:"votes"`
	VirtualLastUpdate     string                     `json:"virtual_last_update"`
	VirtualPosition       string                     `json:"virtual_position"`
	VirtualScheduledTime  string                     `json:"virtual_scheduled_time"`
	TotalMissed           *types.Int                 `json:"total_missed"`
	LastAslot             *types.Int                 `json:"last_aslot"`
	LastConfirmedBlockNum *types.Int                 `json:"last_confirmed_block_num"`
	PowWorker             *types.Int                 `json:"pow_worker"`
	SigningKey            string                     `json:"signing_key"`
	Props                 *ChainProperties           `json:"props"`
	SbdExchangeRate       *CurrentMedianHistoryPrice `json:"sbd_exchange_rate"`
	LastSbdExchangeUpdate *types.TimePointSeconds    `json:"last_sbd_exchange_update"`
	LastWork              string                     `json:"last_work"`
	RunningVersion        string                     `json:"running_version"`
	HardforkVersionVote   string                     `json:"hardfork_version_vote"`
	HardforkTimeVote      *types.TimePointSeconds    `json:"hardfork_time_vote"`
}
