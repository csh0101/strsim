package strsim

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	test string
	need string
}

func Test_ModifyString(t *testing.T) {
	var o option

	o.ignore |= ignoreCase
	o.ignore |= ignoreSpace
	for _, v := range []testCase{
		{
			test: "hello world",
			need: "helloworld",
		},
	} {
		modifyString(&o, &v.test)
		assert.Equal(t, v.test, v.need)
	}
}
