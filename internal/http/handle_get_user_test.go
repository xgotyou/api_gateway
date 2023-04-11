package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xgotyou/api_gateway/internal/dtos"
)

func TestGetUserHandler(t *testing.T) {
	usMock := new(userServiceMock)
	router := SetupRouter(usMock)
	req, _ := http.NewRequest("GET", "/v1/users/10", strings.NewReader(""))
	w := httptest.NewRecorder()
	usMock.On("GetUser", 10).Return(validUserDTO(10), nil)

	router.ServeHTTP(w, req)

	usMock.AssertExpectations(t)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"id":10,"firstName":"John","lastName":"Smith","birthDate":"1985-10-18T03:00:00+00:00","role":"Manager"}`, w.Body.String())
}

func TestGetUserHandlerWithNonIntId(t *testing.T) {
	router := SetupRouter(new(userServiceMock))
	req, _ := http.NewRequest("GET", "/v1/users/alex", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"errors": ["Id must be an integer"]}`, w.Body.String())
}

func validUserDTO(id int) *dtos.User {
	bd := time.Date(1985, time.October, 18, 3, 0, 0, 0, time.FixedZone("Msk", 3))
	return &dtos.User{Id: id, FirstName: "John", LastName: "Smith", BirthDate: &bd, Role: dtos.Manager}
}
