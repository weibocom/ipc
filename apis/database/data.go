package database

import (
	// Stdlib
	"encoding/json"
	"strconv"
	"strings"

	// RPC
	"github.com/weibocom/steem-rpc/steem/types"
)

type Config struct {
	SteemBlockchainHardforkVersion string `json:"STEEM_BLOCKCHAIN_HARDFORK_VERSION"`
	SteemBlockchainVersion         string `json:"STEEM_BLOCKCHAIN_VERSION"`
	SteemBlockInterval             uint   `json:"STEEM_BLOCK_INTERVAL"`
}

type DynamicGlobalProperties struct {
	Time                     *types.TimePointSeconds `json:"time"`
	TotalPow                 *types.Int              `json:"total_pow"`
	NumPowWitnesses          *types.Int              `json:"num_pow_witnesses"`
	CurrentReserveRatio      *types.Int              `json:"current_reserve_ratio"`
	ID                       *types.ID               `json:"id"`
	CurrentSupply            types.Asset             `json:"current_supply"`
	CurrentSBDSupply         types.Asset             `json:"current_sbd_supply"`
	MaximumBlockSize         *types.Int              `json:"maximum_block_size"`
	RecentSlotsFilled        *types.Int              `json:"recent_slots_filled"`
	CurrentWitness           string                  `json:"current_witness"`
	TotalRewardShares2       *types.Int              `json:"total_reward_shares2"`
	AverageBlockSize         *types.Int              `json:"average_block_size"`
	CurrentAslot             *types.Int              `json:"current_aslot"`
	LastIrreversibleBlockNum uint32                  `json:"last_irreversible_block_num"`
	TotalVestingShares       types.Asset             `json:"total_vesting_shares"`
	TotalVersingFundSteem    types.Asset             `json:"total_vesting_fund_steem"`
	HeadBlockID              string                  `json:"head_block_id"`
	HeadBlockNumber          types.UInt32            `json:"head_block_number"`
	VirtualSupply            types.Asset             `json:"virtual_supply"`
	ConfidentialSupply       types.Asset             `json:"confidential_supply"`
	ConfidentialSBDSupply    types.Asset             `json:"confidential_sbd_supply"`
	TotalRewardFundSteem     types.Asset             `json:"total_reward_fund_steem"`
	TotalActivityFundSteem   string                  `json:"total_activity_fund_steem"`
	TotalActivityFundShares  *types.Int              `json:"total_activity_fund_shares"`
	SBDInterestRate          *types.Int              `json:"sbd_interest_rate"`
	MaxVirtualBandwidth      *types.Int              `json:"max_virtual_bandwidth"`
}

type Block struct {
	Number                uint32                  `json:"-"`
	Timestamp             *types.TimePointSeconds `json:"timestamp"`
	Witness               string                  `json:"witness"`
	WitnessSignature      string                  `json:"witness_signature"`
	TransactionMerkleRoot string                  `json:"transaction_merkle_root"`
	Previous              string                  `json:"previous"`
	Extensions            [][]interface{}         `json:"extensions"`
	Transactions          []*types.Transaction    `json:"transactions"`
}

type Content struct {
	Id                      *types.ID               `json:"id"`
	RootTitle               string                  `json:"root_title"`
	Active                  *types.TimePointSeconds `json:"active"`
	AbsRshares              *types.Int              `json:"abs_rshares"`
	PendingPayoutValue      string                  `json:"pending_payout_value"`
	TotalPendingPayoutValue string                  `json:"total_pending_payout_value"`
	Category                string                  `json:"category"`
	Title                   string                  `json:"title"`
	LastUpdate              *types.TimePointSeconds `json:"last_update"`
	Stats                   string                  `json:"stats"`
	Body                    string                  `json:"body"`
	Created                 *types.TimePointSeconds `json:"created"`
	Replies                 []*Content              `json:"replies"`
	Permlink                string                  `json:"permlink"`
	JsonMetadata            *ContentMetadata        `json:"json_metadata"`
	Children                *types.Int              `json:"children"`
	NetRshares              *types.Int              `json:"net_rshares"`
	URL                     string                  `json:"url"`
	ActiveVotes             []*VoteState            `json:"active_votes"`
	ParentPermlink          string                  `json:"parent_permlink"`
	CashoutTime             *types.TimePointSeconds `json:"cashout_time"`
	TotalPayoutValue        string                  `json:"total_payout_value"`
	ParentAuthor            string                  `json:"parent_author"`
	ChildrenRshares2        *types.Int              `json:"children_rshares2"`
	Author                  string                  `json:"author"`
	Depth                   *types.Int              `json:"depth"`
	TotalVoteWeight         *types.Int              `json:"total_vote_weight"`
}

func (content *Content) IsStory() bool {
	return content.ParentAuthor == ""
}

type ContentMetadata struct {
	Flag  bool
	Value string
	Users []string
	Tags  []string
	Image []string
}

type ContentMetadataRaw struct {
	Users types.StringSlice `json:"users"`
	Tags  types.StringSlice `json:"tags"`
	Image types.StringSlice `json:"image"`
}

func (metadata *ContentMetadata) UnmarshalJSON(data []byte) error {
	unquoted, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	switch unquoted {
	case "true":
		metadata.Flag = true
		return nil
	case "false":
		metadata.Flag = false
		return nil
	}

	if len(unquoted) == 0 {
		return nil
	}

	if unquoted[0] == '"' {
		metadata.Value = unquoted
		return nil
	}

	var raw ContentMetadataRaw
	if err := json.NewDecoder(strings.NewReader(unquoted)).Decode(&raw); err != nil {
		return err
	}

	metadata.Users = raw.Users
	metadata.Tags = raw.Tags
	metadata.Image = raw.Image

	return nil
}

type VoteState struct {
	Voter   string                  `json:"voter"`
	Weight  *types.Int              `json:"weight"`
	Rshares *types.Int              `json:"rshares"`
	Percent *types.Int              `json:"percent"`
	Time    *types.TimePointSeconds `json:"time"`
}
