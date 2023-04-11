package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handlCreateUser(us UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var params CreateUserParams
		if err := c.BindJSON(&params); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		u, err := us.CreateUser(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"errors": []string{"Something went wrong"},
			})
		}
		// u := dtos.User{Id: 1, FirstName: params.FirstName, LastName: params.LastName, Role: dtos.Role(params.Role)}

		c.Header("Location", fmt.Sprintf("%v/%v", c.FullPath(), 1))
		c.JSON(http.StatusCreated, u)
	}
}
