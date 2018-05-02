package steem

import (
	"github.com/weibocom/ipc/config"
	"github.com/weibocom/ipc/keys"
)

var privateKeys = make([][]byte, 0, len(config.GetWIFs()))
var publicKeys = make([][]byte, 0, len(config.GetWIFs()))

func init() {
	for _, v := range config.GetWIFs() {
		w, err := keys.DecodeWIF(v)
		if err != nil {
			panic(err)
		}
		privateKeys = append(privateKeys, w.Serialize())
		publicKeys = append(publicKeys, w.PublicKey().Serialize())
	}
}

func GetPrivateKeys() [][]byte {
	return privateKeys
}

func GetPublicKeys() [][]byte {
	return publicKeys
}
