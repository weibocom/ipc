package service

import (
	"fmt"
	"time"

	"git.intra.weibo.com/platform/go-metrics"
)

var (
	accountGauge = metrics.GetOrRegisterGauge("chain.accounts", nil)
	postGauge    = metrics.GetOrRegisterGauge("chain.posts", nil)
)

func StartIPCMetrics() {
	ticker := time.NewTicker(time.Second)

	for range ticker.C {
		uc, err := UserCount()
		if err == nil {
			accountGauge.Update(uc)
		}

		if err != nil {
			fmt.Println(err)
		}

		pc, err := ipcClient.PostCount()
		if err == nil {
			postGauge.Update(int64(pc))
		}

		if err != nil {
			fmt.Println(err)
		}
	}
}
