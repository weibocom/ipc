package service

import (
	"github.com/jinzhu/gorm"
	"github.com/weibocom/ipc/client"
)

var DB *gorm.DB

var (
	IPCClient client.Client
)
