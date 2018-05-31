package content

import "github.com/yanyiwu/gojieba"

// Extract extracts keywords from string s.
func Extract(s string, topk int) []string {
	x := gojieba.NewJieba()
	defer x.Free()
	return x.Extract(s, topk)
}
