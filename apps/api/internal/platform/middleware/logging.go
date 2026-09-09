package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		gin.DefaultWriter.Write([]byte(
			c.Request.Method + " " + c.Request.URL.Path +
				" " + time.Since(start).String() + "\n",
		))
	}
}
