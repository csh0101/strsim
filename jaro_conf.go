package strsim

import (
	"github.com/csh0101/strsim/similarity"
)

// ngram 是筛子系数需要用的一个值
func Jaro(matchWindow ...int) OptionFunc {
	return OptionFunc(func(o *option) {
		mw := 0
		if len(matchWindow) > 0 {
			mw = matchWindow[0]
		}
		d := &similarity.Jaro{MatchWindow: mw}
		o.cmp = d.CompareUtf8
	})
}
