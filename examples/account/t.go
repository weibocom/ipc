package main

import (
	"encoding/json"
	"fmt"

	"github.com/weibocom/steem-rpc/types"
)

func main() {
	m := make(map[string]int64)
	var kam = types.KeyAuthorityMap(m)
	kam["STM6kbKsZj5kY5QrG8huATPtwfVmZmKzFDfUXz1eEbKYF58LorAxF"] = 3
	j, err := json.Marshal(kam)
	if err != nil {
		fmt.Printf("marshal error:%v\n", err.Error())
		return
	}
	fmt.Printf("%s", j)
}
