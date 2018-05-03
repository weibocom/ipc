package types

import (
	"encoding/json"
	"errors"

	"github.com/weibocom/ipc/encoding"
)

// KeyAuthorityMap  refered to `typedef flat_map< public_key_type, weight_type >                key_authority_map;`
type KeyAuthorityMap map[PublicKey]int64

func NewKeyAuthorityMap() KeyAuthorityMap {
	m := make(map[PublicKey]int64)
	return KeyAuthorityMap(m)
}

func (m KeyAuthorityMap) AddAuthority(key string, weightType int64) {
	m[PublicKey(key)] = weightType
}

func (m KeyAuthorityMap) MarshalJSON() ([]byte, error) {
	l := len(m)
	xs := make([]interface{}, l)
	i := 0
	for k, v := range m {
		xs[i] = []interface{}{k, v}
		i++
	}
	return json.Marshal(xs)
}

func (m *KeyAuthorityMap) UnmarshalJSON(data []byte) error {
	var xs [][]interface{}
	if err := json.Unmarshal(data, &xs); err != nil {
		return err
	}

	var invalid bool
	mp := make(map[PublicKey]int64, len(xs))
loop:
	for _, kv := range xs {
		if len(kv) != 2 {
			invalid = true
			break
		}

		k, ok := kv[0].(string)
		if !ok {
			invalid = true
			break
		}

		var v int64
		switch t := kv[1].(type) {
		case float64:
			v = int64(t)
		case int64:
			v = t
		default:
			invalid = true
			break loop
		}

		mp[PublicKey(k)] = v
	}
	if invalid {
		return errors.New("invalid map encoding")
	}

	*m = mp
	return nil
}

func (m KeyAuthorityMap) Len() int {
	return len(m)
}

func (m KeyAuthorityMap) Marshal(encoder *encoding.Encoder) error {
	enc := encoding.NewRollingEncoder(encoder)
	enc.EncodeUVarint(uint64(m.Len()))
	if m.Len() > 0 {
		for k, v := range m {
			enc.Encode(k)
			enc.Encode(uint16(v))
		}
	}
	return enc.Err()
}
