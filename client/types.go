//go:generate jsonenums -type=Pill
package client

import "strconv"

type ContentType uint8

const (
	ContentPost ContentType = iota + 1
	ContentVideo
)

func (ct ContentType) Value() uint8 {
	return uint8(ct)
}

func ParseContentType(s string) ContentType {
	i, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		return ContentPost
	}
	return ContentType(i)
}

type StoreType uint8

const (
	StoreInDB StoreType = iota + 1
	StoreInRedis
	StoreInIPFS
)

func (st StoreType) Value() uint8 {
	return uint8(st)
}
