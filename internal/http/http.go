package http

import (
	"github.com/gin-gonic/gin"
	"github.com/xgotyou/api_gateway/internal/dtos"
)

type UserService interface {
	GetUser(id int) (*dtos.User, error)
	CreateUser(params CreateUserParams) (*dtos.User, error)
}

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Role      string `json:"role"`
}

func SetupRouter(us UserService) *gin.Engine {
	r := gin.Default()

	g := r.Group("v1")
	{
		g.GET("/users/:id", handleGetUser(us))
		g.POST("/users", handlCreateUser(us))
	}

	return r
}
