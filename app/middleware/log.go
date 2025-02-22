package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LogMiddleware logs each request
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		statusCode := c.Writer.Status()
		log.Printf("%s %s | %d | %s | %s", c.Request.Method, c.Request.URL.Path, statusCode, latency, c.ClientIP())
	}
}
