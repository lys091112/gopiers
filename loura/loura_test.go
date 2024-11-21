package loura

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func TestStart(t *testing.T) {
	bearerToken := "Bearer token"
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	pprof.Register(r)
	adminGroup := r.Group("/admin", func(c *gin.Context) {
		if c.Request.Header.Get("Authorization") != bearerToken {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	})
	pprof.RouteRegister(adminGroup, "pprof")

	req, _ := http.NewRequest(http.MethodGet, "/debug/pprof/", nil)

	rw := httptest.NewRecorder()
	r.ServeHTTP(rw, req)

	if excepted, got := http.StatusOK, rw.Code; excepted != got {
		t.Errorf("excepted:%d,got:%d", excepted, got)
	}

	req, _ = http.NewRequest(http.MethodGet, "/admin/pprof/", nil)
	rw = httptest.NewRecorder()
	r.ServeHTTP(rw, req)

	if expected, got := http.StatusForbidden, rw.Code; expected != got {
		t.Errorf("expected: %d, got: %d", expected, got)
	}

	req, _ = http.NewRequest(http.MethodGet, "/admin/pprof/", nil)
	req.Header.Set("Authorization", bearerToken)
	rw = httptest.NewRecorder()
	r.ServeHTTP(rw, req)

	if expected, got := http.StatusOK, rw.Code; expected != got {
		t.Errorf("expected: %d, got: %d", expected, got)
	}
}
