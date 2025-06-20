package middleware

import (
	"backend-api/utils/logger"
	"github.com/gin-gonic/gin"
	"time"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		status := c.Writer.Status()

		logger.Logger.Printf("[GIN] %d | %s | %s | %v\n",
			status,
			c.Request.Method,
			c.Request.URL.Path,
			duration,
		)
	}
}
