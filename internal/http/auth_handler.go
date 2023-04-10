package http

import (
	"fmt"
	"log"
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

		if err != nil || !token.Valid {
			c.String(http.StatusUnauthorized, err.Error())
			return
		}

		if id, ok := token.Claims.(jwt.MapClaims)["id"].(float64); ok {
			c.Set("user_id", int(id))
		} else {
			log.Println("warn: claim 'id' not found in the token")
		}
		if role, ok := token.Claims.(jwt.MapClaims)["role"].(string); ok {
			c.Set("user_role", role)
		} else {
			log.Println("warn: claim 'role' not found in the token")
		}

		c.Next()
	}
}
