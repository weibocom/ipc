package condenser

import "github.com/weibocom/ipc/steem/types"

type AccountKeys struct {
	WeightThreshold *types.Int    `json:"weight_threshold"`
	AccountAuths    []interface{} `json:"account_auths"`
	KeyAuths        []interface{} `json:"key_auths"`
}
type Account struct {
	ID                *types.Int              `json:"id"`
	Name              string                  `json:"name"`
	Owner             *AccountKeys            `json:"owner"`
	Active            *AccountKeys            `json:"active"`
	Posting           *AccountKeys            `json:"posting"`
	MemoKey           string                  `json:"memo_key"`
	JSONMetadata      string                  `json:"json_metadata"`
	Proxy             string                  `json:"proxy"`
	LastOwnerUpdate   *types.TimePointSeconds `json:"last_owner_update"`
	LastAccountUpdate *types.TimePointSeconds `json:"last_account_update"`
	Created           *types.TimePointSeconds `json:"created"`
	Mined             bool                    `json:"mined"`
	OwnerChallenged   bool                    `json:"owner_challenged"`
	ActiveChallenged  bool                    `json:"active_challenged"`
	//LastOwnerProved               *types.Time   `json:"last_owner_proved"`
	//LastActiveProved              *types.Time   `json:"last_active_proved"`
	RecoveryAccount               string                  `json:"recovery_account"`
	LastAccountRecovery           *types.TimePointSeconds `json:"last_account_recovery"`
	ResetAccount                  string                  `json:"reset_account"`
	CommentCount                  *types.Int              `json:"comment_count"`
	LifetimeVoteCount             *types.Int              `json:"lifetime_vote_count"`
	PostCount                     *types.Int              `json:"post_count"`
	CanVote                       bool                    `json:"can_vote"`
	VotingPower                   int                     `json:"voting_power"`
	LastVoteTime                  *types.TimePointSeconds `json:"last_vote_time"`
	Balance                       string                  `json:"balance"`
	SavingsBalance                string                  `json:"savings_balance"`
	SbdBalance                    string                  `json:"sbd_balance"`
	SbdSeconds                    string                  `json:"sbd_seconds"`
	SbdSecondsLastUpdate          *types.TimePointSeconds `json:"sbd_seconds_last_update"`
	SbdLastInterestPayment        *types.TimePointSeconds `json:"sbd_last_interest_payment"`
	SavingsSbdBalance             string                  `json:"savings_sbd_balance"`
	SavingsSbdSeconds             string                  `json:"savings_sbd_seconds"`
	SavingsSbdSecondsLastUpdate   *types.TimePointSeconds `json:"savings_sbd_seconds_last_update"`
	SavingsSbdLastInterestPayment *types.TimePointSeconds `json:"savings_sbd_last_interest_payment"`
	SavingsWithdrawRequests       *types.Int              `json:"savings_withdraw_requests"`
	VestingShares                 string                  `json:"vesting_shares"`
	VestingWithdrawRate           string                  `json:"vesting_withdraw_rate"`
	NextVestingWithdrawal         *types.TimePointSeconds `json:"next_vesting_withdrawal"`
	Withdrawn                     *types.Int              `json:"withdrawn"`
	ToWithdraw                    *types.Int              `json:"to_withdraw"`
	WithdrawRoutes                *types.Int              `json:"withdraw_routes"`
	CurationRewards               *types.Int              `json:"curation_rewards"`
	PostingRewards                *types.Int              `json:"posting_rewards"`
	ProxiedVsfVotes               []*types.Int            `json:"proxied_vsf_votes"`
	WitnessesVotedFor             *types.Int              `json:"witnesses_voted_for"`
	AverageBandwidth              *types.Int              `json:"average_bandwidth"`
	LifetimeBandwidth             *types.Int              `json:"lifetime_bandwidth"`
	LastBandwidthUpdate           *types.TimePointSeconds `json:"last_bandwidth_update"`
	AverageMarketBandwidth        *types.Int              `json:"average_market_bandwidth"`
	LastMarketBandwidthUpdate     *types.TimePointSeconds `json:"last_market_bandwidth_update"`
	LastPost                      *types.TimePointSeconds `json:"last_post"`
	LastRootPost                  *types.TimePointSeconds `json:"last_root_post"`
	PostBandwidth                 *types.Int              `json:"post_bandwidth"`
	NewAverageBandwidth           string                  `json:"new_average_bandwidth"`
	NewAverageMarketBandwidth     *types.Int64            `json:"new_average_market_bandwidth"`
	VestingBalance                string                  `json:"vesting_balance"`
	Reputation                    *types.Int64            `json:"reputation"`
	TransferHistory               []interface{}           `json:"transfer_history"`
	MarketHistory                 []interface{}           `json:"market_history"`
	PostHistory                   []interface{}           `json:"post_history"`
	VoteHistory                   []interface{}           `json:"vote_history"`
	OtherHistory                  []interface{}           `json:"other_history"`
	WitnessVotes                  []string                `json:"witness_votes"`
	TagsUsage                     []interface{}           `json:"tags_usage"`
	GuestBloggers                 []interface{}           `json:"guest_bloggers"`
	BlogCategory                  interface{}             `json:"blog_category"`
}
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

type FeedEntry struct {
	Author   string `json:"string"`
	Permlink string `json:"permlink"`
	EntryID  uint32 `json:"entry_id"`
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
	JsonMetadata            string                  `json:"json_metadata"`
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

type VoteState struct {
	Voter   string                  `json:"voter"`
	Weight  *types.Int              `json:"weight"`
	Rshares *types.Int              `json:"rshares"`
	Percent *types.Int              `json:"percent"`
	Time    *types.TimePointSeconds `json:"time"`
}
