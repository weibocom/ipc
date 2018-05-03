package util

import "unsafe"

// Bytes2String converts []byte to string.
func Bytes2String(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

// String2Bytes converts string to []byte.
func String2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
