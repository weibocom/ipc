package steem

const (
	STEEM_ADDRESS_PREFIX = "STM"
)

var wifs = []string{
	"5JzpcbsNCu6Hpad1TYmudH4rj1A22SW9Zhb1ofBGHRZSp5poqAX",
}

// var privateKeys = make([][]byte, 0, len(wifs))
// var publicKeys = make([][]byte, 0, len(wifs))

// func init() {
// 	for _, v := range wifs {
// 		privKey, err := wif.Decode(v)
// 		if err != nil {
// 			panic(err)
// 		}
// 		privateKeys = append(privateKeys, privKey)
// 	}

// 	for _, v := range wifs {
// 		pubKey, err := wif.GetPublicKey(v)
// 		if err != nil {
// 			panic(err)
// 		}
// 		publicKeys = append(publicKeys, pubKey)
// 	}
// }

// func GetWifs() []string {
// 	return wifs
// }

// func GetPrivateKeys() [][]byte {
// 	return privateKeys
// }

// func GetPublicKeys() [][]byte {
// 	return publicKeys
// }
