package content

import (
	"path"

	"github.com/yanyiwu/gojieba"
)

var x = gojieba.NewJieba()

// Extract extracts keywords from string s.
func Extract(s string, topk int) []string {
	return x.Extract(s, topk)
}

func ConfigGojieba(d string) {
	if x != nil {
		x.Free()
	}

	dpath := path.Join(d, "jieba.dict.utf8")
	hpath := path.Join(d, "hmm_model.utf8")
	upath := path.Join(d, "user.dict.utf8")
	ipath := path.Join(d, "idf.utf8")
	spath := path.Join(d, "stop_words.utf8")

	x = gojieba.NewJieba(dpath, hpath, upath, ipath, spath)
}

func CleanGojieba() {
	x.Free()
}
