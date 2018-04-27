package types

import (
	// Stdlib

	"time"

	"github.com/weibocom/steem-rpc/encoding/transaction"
)

const Layout = `"2006-01-02T15:04:05"`

var offset int

func init() {
	_, offset = time.Now().Zone()
}

// TimePointSeconds UTC时区 steem使用的是UTC时区
type TimePointSeconds struct {
	*time.Time
}

func NewTimePointSeconds(t time.Time) *TimePointSeconds {
	return &TimePointSeconds{
		&t,
	}
}

func (t *TimePointSeconds) MarshalJSON() ([]byte, error) {
	tm := time.Unix(t.Unix()-int64(offset), 0)
	return []byte(tm.Format(Layout)), nil
}

func (t *TimePointSeconds) UnmarshalJSON(data []byte) error {
	parsed, err := time.ParseInLocation(Layout, string(data), time.UTC)
	if err != nil {
		return err
	}
	t.Time = &parsed
	return nil
}

func (t *TimePointSeconds) MarshalTransaction(encoder *transaction.Encoder) error {
	return encoder.Encode(uint32(t.Time.Unix()))
}
