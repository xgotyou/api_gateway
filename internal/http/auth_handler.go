package http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"

	"github.com/gin-gonic/gin"
)

func JwtAuthHandler(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		// Check if the Authorization header is present and valid
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.String(http.StatusUnauthorized, "Invalid Authorization header")
			return
		}

		// Extract the JWT token from the Authorization header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse and validate the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secret), nil
		})

		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}

		if !token.Valid {
			c.String(http.StatusUnauthorized, "Invalid token")
			return
		}

		c.Next()
	}
}
