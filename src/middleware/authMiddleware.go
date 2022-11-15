package middleware

import (
	helper "app/src/helpers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		incomingToken := c.Request.Header.Get("Authorization")
		if incomingToken == "" {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		token := strings.Split(incomingToken, "Bearer ")
		claims, err := helper.ValidateToken(token[1])
		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("username", claims.Username)
		c.Set("id", claims.Uid)
		session := sessions.Default(c)
		session.Set("username", claims.Username)
		session.Set("id", claims.Uid)
		c.Next()
	}
}
