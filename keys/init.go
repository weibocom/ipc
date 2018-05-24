package keys

import "github.com/weibocom/ipc/config"

var privateKeys = make([][]byte, 0, len(config.GetWIFs()))
var publicKeys = make([][]byte, 0, len(config.GetWIFs()))

func init() {
	for _, v := range config.GetWIFs() {
		w, err := DecodeWIF(v)
		if err != nil {
			panic(err)
		}
		privateKeys = append(privateKeys, w.Serialize())
		publicKeys = append(publicKeys, w.PublicKey().Serialize())
	}
}

// InitKeys init those keys with given wifs.
// It is useful when apps start and use different keys with initminer.
func InitKeys(wifs ...string) {
	var priKeys = make([][]byte, 0, len(wifs))
	var pubKeys = make([][]byte, 0, len(wifs))

	for _, v := range wifs {
		w, err := DecodeWIF(v)
		if err != nil {
			panic(err)
		}
		priKeys = append(privateKeys, w.Serialize())
		pubKeys = append(publicKeys, w.PublicKey().Serialize())
	}

	privateKeys = priKeys
	publicKeys = pubKeys
}

func GetPrivateKeys() [][]byte {
	return privateKeys
}

func GetPublicKeys() [][]byte {
	return publicKeys
}
