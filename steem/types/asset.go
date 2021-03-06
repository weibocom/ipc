package types

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/weibocom/ipc/encoding"
)

type Asset []interface{}

func NewSteemAsset(fee int64) Asset {
	asset := Asset{
		fee,
		3,
		"@@000000021",
	}
	return asset
}

// 一共16个字节。8个字节存储fee。1个字节存储类型。7个字节存储SYMBOL
func (a Asset) Marshal(encoder *encoding.Encoder) error {
	fee, err := convertToInt64(a[0])
	if err != nil {
		return err
	}

	enc := encoding.NewRollingEncoder(encoder)
	enc.Encode(fee) // 8 bytes
	// symbol 一共8个字节 第一个字节3表示类型：是STEEM、SBD等。
	symbol := []byte{3, 'S', 'T', 'E', 'E', 'M', 0, 0}
	enc.Encode(symbol)

	return enc.Err()
}

func (a *Asset) UnmarshalJSON(data []byte) error {

	s := string(data)
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")
	ss := strings.Split(s, ",")

	// fee returned by GetAccountHistory is like `0.001 STEEM`, not like CreateAccount `["1", 3, "@@000000021"]``
	if len(ss) == 1 {
		s = strings.Trim(s, "\"")
		ss = strings.Split(s, " ")
		ss[1] = "3"
		ss = append(ss, "@@000000021")
	}
	if len(ss) < 3 {
		return fmt.Errorf("wrong fee format: %s, %s", s, ss)
	}

	fee, err := strconv.ParseFloat(strings.Trim(ss[0], "\""), 32)
	if err != nil {
		return err
	}

	sym, err := strconv.Atoi(ss[1])
	if err != nil {
		return err
	}

	*a = Asset{
		fee,
		sym,
		ss[2],
	}

	return nil
}

func convertToInt64(n interface{}) (int64, error) {
	switch v := n.(type) {
	case int:
		return int64(v), nil
	case float32:
		return int64(v), nil
	case float64:
		return int64(v), nil
	case int64:
		return int64(v), nil
	default:
		s := fmt.Sprintf("%v", n)
		i, err := strconv.Atoi(s)
		return int64(i), err
	}
}
