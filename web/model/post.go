package model

import "github.com/weibocom/ipc/model"

type Post struct {
	*model.Post
	Similarity float64 `json:"similarity,omitempty"`
}
