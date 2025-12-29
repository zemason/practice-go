package middleware

import (
	"fmt"
	"time"
	
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start time
		start := time.Now()
		
		// Process request
		c.Next()
		
		// End time
		end := time.Now()
		latency := end.Sub(start)
		
		// Log information
		fmt.Printf("[GIN] %s | %3d | %13v | %15s | %-7s %s\n",
			end.Format("2006/01/02 - 15:04:05"),
			c.Writer.Status(),
			latency,
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
		)
	}
}