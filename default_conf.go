package strsim

import (
	"github.com/csh0101/strsim/similarity"
)

func Default() OptionFunc {
	return OptionFunc(func(o *option) {
		if o.cmp == nil {
			l := similarity.EditDistance{}
			o.cmp = l.CompareUtf8
			if o.ascii {
				o.cmp = l.CompareAscii
			}
		}
	})
}
