package io

import (
	"io"
)

func DebugWriter(w io.Writer) io.Writer {
	return &debug_writer{
		w,
	}
}

type debug_writer struct {
	io.Writer
}

func (d debug_writer) Write(p []byte) (n int, err error) {
	return d.Writer.Write(p)
}
