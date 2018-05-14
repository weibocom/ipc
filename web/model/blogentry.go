package model

import "time"

type BlogEntry struct {
	Author   string    `json:"author,omitempty"`
	Blog     string    `json:"blog,omitempty"`
	Permlink string    `json:"permlink,omitempty"`
	EntryID  uint32    `json:"entry_id,omitempty"`
	ReblogOn time.Time `json:"reblog_on,omitempty"`
}
