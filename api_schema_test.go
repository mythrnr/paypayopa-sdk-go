package paypayopa_test

import (
	"net/http"
	"testing"

	"github.com/mythrnr/paypayopa-sdk-go"
	"github.com/stretchr/testify/assert"
)

func Test_ResultInfo_Success(t *testing.T) {
	t.Parallel()

	tests := []struct {
		statusCode int
		want       bool
	}{
		{statusCode: http.StatusOK, want: true},
		{statusCode: http.StatusBadRequest, want: false},
		{statusCode: http.StatusInternalServerError, want: false},
	}

	for _, tt := range tests {
		info := new(paypayopa.ResultInfo)
		info.StatusCode = tt.statusCode

		assert.Equal(t, tt.want, info.Success())
	}
}
