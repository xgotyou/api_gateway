package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const secret = "known"

// Token encoded using HS256, payload: { "id": 1, "role": "Manager" }
const token = "eyJhbGciOiJIUzI1NiJ9.eyJpZCI6MSwicm9sZSI6Ik1hbmFnZXIifQ.rYZG0Bl1lb9T_qZQaIS5jPSScFdRa6o5QlnwlSDiO_M"

func TestAuthHandlerValidToken(t *testing.T) {
	r := gin.New()
	r.Use(JwtAuthHandler(secret))

	req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))

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
