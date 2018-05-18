package model

import "time"

type User struct {
	ID         int64     `json:"id,omitempty"`
	Company    string    `json:"company,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	PostCount  int       `json:"post_count"`
	PrivateKey string    `json:"private_key,omitempty"`
	PublicKey  string    `json:"public_key,omitempty"`
}
