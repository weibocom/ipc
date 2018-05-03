package content

import (
	"strings"

	"github.com/rfguri/bowsim"
	"github.com/yanyiwu/gojieba"
)

func Similarity(a, b string) float64 {
	x := gojieba.NewJieba()
	defer x.Free()

	words1 := x.Cut(a, true)
	words2 := x.Cut(b, true)

	return bowsim.Get(strings.Join(words1, " "), strings.Join(words2, " "))
}
