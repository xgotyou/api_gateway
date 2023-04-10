package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	secret = "known"

	// Token encoded using HS256, payload: { "id": 1, "role": "Manager" }
	validToken  = "eyJhbGciOiJIUzI1NiJ9.eyJpZCI6MSwicm9sZSI6Ik1hbmFnZXIifQ.rYZG0Bl1lb9T_qZQaIS5jPSScFdRa6o5QlnwlSDiO_M"
	noIdToken   = "eyJhbGciOiJIUzI1NiJ9.eyJyb2xlIjoiTWFuYWdlciJ9.yl01v0cZiEmbEsOsZhvjv1wZhN_HjQZjSLgrm3fHmWY"
	noRoleToken = "eyJhbGciOiJIUzI1NiJ9.eyJpZCI6MX0.OXzUmbUDk0AAiGy_lWkw0gsAlRM9XIbvip7AhFxtkkU"
)

func TestAuthHandlerValidToken(t *testing.T) {
	r := gin.New()
	r.Use(JwtAuthHandler(secret))

	req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", validToken))

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
}

func TestAuthHandlerNoAuthHeader(t *testing.T) {
	r := gin.New()
	r.Use(JwtAuthHandler(secret))

	req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
	// Note that "Authorization" header is not set

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthHandlerNoToken(t *testing.T) {
	r := gin.New()
	r.Use(JwtAuthHandler(secret))

	req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", "broken-token"))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthHandlerIncorrectAuthHeader(t *testing.T) {
	r := gin.New()
	r.Use(JwtAuthHandler(secret))

	req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
	req.Header.Add("Authorization", fmt.Sprintf("Basic %v", "YWRtaW46cm9vdA==\n"))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthHandlerSetsIdAndRoleToContext(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", validToken))
	c := new(gin.Context)
	c.Request = req

	JwtAuthHandler(secret)(c)

	assert.Equal(t, req, c.Request)
	assert.Equal(t, 1, c.GetInt("user_id"))
	assert.Equal(t, "Manager", c.GetString("user_role"))
}

func TestAuthHandlerHandlesTokenWithoutId(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", noIdToken))
	c := new(gin.Context)
	c.Request = req

	JwtAuthHandler(secret)(c)

	assert.Equal(t, req, c.Request)
	_, found := c.Get("user_id")
	assert.False(t, found)
	assert.Equal(t, "Manager", c.GetString("user_role"))
}

func TestAuthHandlerHandlesTokenWithoutRole(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", noRoleToken))
	c := new(gin.Context)
	c.Request = req

	JwtAuthHandler(secret)(c)

	assert.Equal(t, req, c.Request)
	_, found := c.Get("user_role")
	assert.False(t, found)
	assert.Equal(t, 1, c.GetInt("user_id"))
}
