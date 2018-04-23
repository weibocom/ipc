package config

// const (
// 	steemAddressPrefix = "STM"
// 	chainID            = "2ac122bd70a2f74d6f761c331f4c4da1028b783cc185d23bf5449ac5af49e198"
// 	testChainID        = "18dcf0a285365fc58b71f18b3d3fec954aa0c141c44e4e5cb4cf777b9eab274e"
// )

// var wifs = []string{"5JzpcbsNCu6Hpad1TYmudH4rj1A22SW9Zhb1ofBGHRZSp5poqAX"}

var defaultConfig = SteemConfig{
	Wifs:               []string{"5JzpcbsNCu6Hpad1TYmudH4rj1A22SW9Zhb1ofBGHRZSp5poqAX"},
	ChainID:            "2ac122bd70a2f74d6f761c331f4c4da1028b783cc185d23bf5449ac5af49e198",
	SteemAddressPrefix: "STM",
}

type SteemConfig struct {
	Wifs               []string
	ChainID            string
	SteemAddressPrefix string
}

func SetConfig(conf SteemConfig) {
	defaultConfig = conf
}

func GetWIFs() []string {
	return defaultConfig.Wifs
}

func GetChainID() string {
	return defaultConfig.ChainID
}

func GetAddressPrefix() string {
	return defaultConfig.SteemAddressPrefix
}
