package model

import "github.com/weibocom/ipc/model"

type Post struct {
	*model.Post
	Similarity string `json:"similarity"`
}
