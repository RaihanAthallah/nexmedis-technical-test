package middleware

import (
	"net/http"
	"nexmedis-technical-test/app/auth"
	"nexmedis-technical-test/app/helper"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a middleware function that validates JWT access tokens
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
			c.Abort()
			return
		}

		// Check if token is in Bearer format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization format"})
			c.Abort()
			return
		}
		tokenString := parts[1]

		// Validate the token
		claims, err := auth.ValidateAccessToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		// decrypt ID
		id, err := helper.DecryptString(claims.ID)
		intId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		// Set claims in the context for use in handlers
		c.Set("username", claims.Username)
		c.Set("email", claims.Email)
		c.Set("user_id", intId)
		// Proceed to the next middleware/handler
		c.Next()
	}
}
