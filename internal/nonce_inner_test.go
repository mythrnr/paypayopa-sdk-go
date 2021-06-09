package internal

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ptn1   = regexp.MustCompile("[0-9a-zA-Z]")
	ptn10  = regexp.MustCompile("[0-9a-zA-Z]{10}")
	ptn100 = regexp.MustCompile("[0-9a-zA-Z]{100}")
)

func Test_Nonce(t *testing.T) {
	t.Parallel()

	tests := []struct {
		num  uint
		rgx  *regexp.Regexp
		want bool
	}{{
		num:  0,
		rgx:  ptn1,
		want: false,
	}, {
		num:  1,
		rgx:  ptn1,
		want: true,
	}, {
		num:  10,
		rgx:  ptn10,
		want: true,
	}, {
		num:  100,
		rgx:  ptn100,
		want: true,
	}}

	for _, tt := range tests {
		t.Log(tt.rgx.String(), tt.num)

		assert.Equal(t,
			tt.want,
			tt.rgx.Match([]byte(nonce(tt.num))),
		)
	}
}
