package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestChangeDecimalToBinary(t *testing.T) {
	tests := []struct {
		name       string
		initRouter func() (*gin.Context, *httptest.ResponseRecorder)
	}{
		{
			name: "flow error params",
			initRouter: func() (*gin.Context, *httptest.ResponseRecorder) {
				h := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(h)
				c.Request, _ = http.NewRequest(http.MethodGet, "localhost:8085/binary?query=io", nil)
				return c, h
			},
		},
		{
			name: "flow success normal",
			initRouter: func() (*gin.Context, *httptest.ResponseRecorder) {
				h := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(h)
				c.Request, _ = http.NewRequest(http.MethodGet, "localhost:8085/binary?query=110110", nil)
				return c, h
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := tt.initRouter()
			ChangeDecimalToBinary(c)
		})
	}
}
