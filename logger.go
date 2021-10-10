package ginmiddleware

import (
	"time"

	"github.com/designsbysm/logger/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()

		c.Next()

		if !viper.GetBool("gin.release") {
			return
		}

		latency := time.Since(now)
		logger.Write(logger.LevelInfo, c.Writer.Status(), c.Request.Method, c.Request.URL, latency.Round(time.Millisecond))
	}
}
