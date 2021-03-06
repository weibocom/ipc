package model

import (
	"time"
)

type Account struct {
	Name      string    `gorm:"COLUMN:name;PRIMARY_KEY;TYPE:VARCHAR(64);NOT NULL" json:"name,omitempty"`
	Company   string    `gorm:"COLUMN:company;TYPE:VARCHAR(64);NOT NULL" json:"company,omitempty"`
	WIF       string    `gorm:"COLUMN:wif;TYPE:VARCHAR(128);NOT NULL" json:"wif,omitempty"`
	CreatedAt time.Time `gorm:"COLUMN:created_at;" json:"created_at,omitempty"`
}

type DNA []byte

func (dna DNA) String() string {
	return string(dna)
}

func (dna DNA) Bytes() []byte {
	return []byte(dna)
}

type Member struct {
	Name       string    `gorm:"COLUMN:name;PRIMARY_KEY;TYPE:VARCHAR(64);NOT NULL"  json:"name,omitempty"`
	ID         int64     `gorm:"COLUMN:id;NOT NULL;unique" json:"id,omitempty"`
	Company    string    `gorm:"COLUMN:company"  json:"company,omitempty"`
	SigningKey string    `gorm:"COLUMN:signing_key;TYPE:VARCHAR(128);NOT NULL" json:"signing_key,omitempty"`
	Wif        string    `gorm:"COLUMN:wif;TYPE:VARCHAR(128);NOT NULL" json:"wif,omitempty"`
	CreatedAt  time.Time `gorm:"COLUMN:created_at;" json:"created_at,omitempty"`
}

type Content []byte

type Post struct {
	MSGID       int64     `gorm:"COLUMN:mid;NOT NULL" json:"mid,omitempty"`
	DNA         string    `gorm:"COLUMN:dna;index:idx_dna;TYPE:VARCHAR(255);NOT NULL" json:"dna,omitempty"`
	Author      string    `gorm:"COLUMN:author;TYPE:VARCHAR(64);NOT NULL;index:idx_author" json:"author,omitempty"`
	Content     string    `gorm:"COLUMN:content;TYPE:TEXT;NOT NULL" json:"content,omitempty"`
	ContentType uint8     `gorm:"COLUMN:content_type;TYPE:TINYINT" json:"content_type,omitempty"`
	StoreType   uint8     `gorm:"COLUMN:store_type;TYPE:TINYINT" json:"store_type,omitempty"`
	Keywords    string    `gorm:"COLUMN:keywords;TYPE:VARCHAR(256);index:idx_keywords" json:"keywords,omitempty"`
	Digest      string    `gorm:"COLUMN:digest;TYPE:VARCHAR(64);NOT NULL" json:"digest,omitempty"`
	CreatedAt   time.Time `gorm:"COLUMN:created_at;NOT NULL" json:"created_at,omitempty"`
}
