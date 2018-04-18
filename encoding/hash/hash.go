package hash

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"

	"github.com/weibocom/steem-rpc/encoding/transaction"
)

// Hash256 1. 用Encoder将v序列化； 2. 进行hash256处理；3. 输出hex
func Hash256(v interface{}) (string, error) {
	var b bytes.Buffer
	encoder := transaction.NewEncoder(&b)
	if err := encoder.Encode(v); err != nil {
		return "", err
	}
	hash := sha256.Sum256(b.Bytes())
	return hex.EncodeToString(hash[:]), nil
}
