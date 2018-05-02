package config

import "github.com/weibocom/ipc/keys"

var defaultConfig = IPCConfig{
	Wifs:          []string{"5JzpcbsNCu6Hpad1TYmudH4rj1A22SW9Zhb1ofBGHRZSp5poqAX"},
	ChainID:       "2ac122bd70a2f74d6f761c331f4c4da1028b783cc185d23bf5449ac5af49e198",
	AddressPrefix: "STM",
}

var privateKeys = make([][]byte, 0, len(GetWIFs()))
var publicKeys = make([][]byte, 0, len(GetWIFs()))

func init() {
	for _, v := range GetWIFs() {
		w, err := keys.DecodeWIF(v)
		if err != nil {
			panic(err)
		}
		privateKeys = append(privateKeys, w.Serialize())
		publicKeys = append(publicKeys, w.PublicKey().Serialize())
	}
}

type IPCConfig struct {
	Wifs          []string
	ChainID       string
	AddressPrefix string
}

func SetConfig(conf IPCConfig) {
	defaultConfig = conf
}

func GetWIFs() []string {
	return defaultConfig.Wifs
}

func GetChainID() string {
	return defaultConfig.ChainID
}

func GetAddressPrefix() string {
	return defaultConfig.AddressPrefix
}

func GetCreator() string {
	return "initminer"
}

func GetPrivateKeys() [][]byte {
	return privateKeys
}

func GetPublicKeys() [][]byte {
	return publicKeys
}

func GetCreateAccountFee() int64 {
	return 1
}
