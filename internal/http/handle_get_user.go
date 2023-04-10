package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func handleGetUser(us UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errors": []string{"Id must be an integer"},
			})
			return
		}

		u, err := us.GetUser(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"errors": []string{"Something went wrong"},
			})
		}

		c.JSON(http.StatusOK, u)
	}
}
