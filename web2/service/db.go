package service

import (
	"github.com/jinzhu/gorm"
	"github.com/weibocom/ipc/client"
)

var (
	db        *gorm.DB
	ipcClient client.Client
)

func SetIPCClient(c client.Client) {
	ipcClient = c
}

func SetDB(db *gorm.DB) {
	db = db
}
