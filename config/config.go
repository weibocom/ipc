package config

var defaultConfig = IPCConfig{
	Wifs:          []string{"5JzpcbsNCu6Hpad1TYmudH4rj1A22SW9Zhb1ofBGHRZSp5poqAX"},
	ChainID:       "2ac122bd70a2f74d6f761c331f4c4da1028b783cc185d23bf5449ac5af49e198",
	AddressPrefix: "STM",
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
