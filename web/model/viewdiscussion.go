package model

import "time"

type ViewDiscussion struct {
	AuthorName      string    `json:"authorName,omitempty"`
	BlockCreateTime time.Time `json:"blockCreateTime,omitempty"`
	Meta            string    `json:"meta,omitempty"`
	Url             string    `json:"url,omitempty"`
	Content         string    `json:"content,omitempty"`
	BlockTitle      string    `json:"blockTitle,omitempty"`
}
