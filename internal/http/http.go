package http

import (
	"github.com/gin-gonic/gin"
	"github.com/xgotyou/api_gateway/internal/dtos"
)

type UserService interface {
	GetUser(id int) (*dtos.User, error)
}

func SetupRouter(us UserService) *gin.Engine {
	r := gin.Default()

	g := r.Group("v1")
	{
		g.GET("/users/:id", handleGetUser(us))
	}

	return r
}
