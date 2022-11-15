package middleware

import (
	"github.com/gin-gonic/gin"
)

func Open() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
