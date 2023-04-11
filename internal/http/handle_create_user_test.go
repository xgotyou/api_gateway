package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xgotyou/api_gateway/internal/dtos"
)

func TestCreateUserHandler(t *testing.T) {
	userParams := `{"firstName":"Will","lastName":"Smith","role":"Customer"}`
	usMock := new(userServiceMock)
	router := SetupRouter(usMock)
	req, _ := http.NewRequest(http.MethodPost, "/v1/users", strings.NewReader(userParams))
	req.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()
	usMock.
		On("CreateUser", CreateUserParams{"Will", "Smith", "Customer"}).
		Return(&dtos.User{Id: 1, FirstName: "Will", LastName: "Smith", Role: "Customer"}, nil)

	router.ServeHTTP(w, req)

	usMock.AssertExpectations(t)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "/v1/users/1", w.Header().Get("Location"))

	expJSON := `{"id":1,"firstName":"Will","lastName":"Smith","role":"Customer"}`
	assert.JSONEq(t, expJSON, w.Body.String())
}
