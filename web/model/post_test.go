package model

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/weibocom/ipc/model"
)

func TestPost(t *testing.T) {
	p := &model.Post{
		MSGID:     100,
		DNA:       "DNAString",
		Author:    "wushang",
		Title:     "23123123",
		Content:   "contentstr",
		Keywords:  "sadas,sadas,sadsa",
		Digest:    "asdasdsadsadasdsfd",
		CreatedAt: time.Now(),
	}

	pp := &Post{}
	pp.Post = p
	pp.Similarity = 98.72

	buf, _ := json.Marshal(pp)
	t.Logf("post: %s", buf)
}
