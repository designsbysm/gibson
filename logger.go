package gibson

import (
	"time"

	"github.com/designsbysm/logger/v2"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()

		c.Next()

		if gin.IsDebugging() {
			return
		}

		latency := time.Since(now)
		logger.Info(c.Writer.Status(), c.Request.Method, c.Request.URL, latency.Round(time.Millisecond))
	}
}
