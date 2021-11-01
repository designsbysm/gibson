package gibson

import (
	"strings"

	"github.com/designsbysm/timber/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Error() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if gin.IsDebugging() {
			return
		}

		if len(c.Errors.Errors()) > 0 {
			message := strings.Join(c.Errors.Errors(), "; ")
			message = strings.Replace(message, "ERROR: ", "", -1)

			timber.Error(message)

			c.JSON(0, gin.H{
				"error": message,
			})
		}
	}
}
