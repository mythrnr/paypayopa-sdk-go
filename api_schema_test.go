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
		info *paypayopa.ResultInfo
		want bool
	}{{
		info: &paypayopa.ResultInfo{
			StatusCode: http.StatusOK,
		},
		want: true,
	}, {
		info: &paypayopa.ResultInfo{
			StatusCode: http.StatusBadRequest,
		},
		want: false,
	}, {
		info: &paypayopa.ResultInfo{
			StatusCode: http.StatusInternalServerError,
		},
		want: false,
	}}

	for _, tt := range tests {
		assert.Equal(t, tt.info.Success(), tt.want)
	}
}
