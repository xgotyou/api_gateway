package http

import (
	"github.com/stretchr/testify/mock"
	"github.com/xgotyou/api_gateway/internal/dtos"
)

type userServiceMock struct {
	mock.Mock
}

func (us *userServiceMock) GetUser(id int) (*dtos.User, error) {
	args := us.Called(id)
	return args.Get(0).(*dtos.User), args.Error(1)
}

func (us *userServiceMock) CreateUser(params CreateUserParams) (*dtos.User, error) {
	args := us.Called(params)
	return args.Get(0).(*dtos.User), args.Error(1)
}
