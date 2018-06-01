package content

import "github.com/yanyiwu/gojieba"

var x = gojieba.NewJieba()

// Extract extracts keywords from string s.
func Extract(s string, topk int) []string {
	return x.Extract(s, topk)
}

func Clean() {
	x.Free()
}
