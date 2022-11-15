package helper

import (
	model "app/src/models"
	"github.com/gin-gonic/gin"
	"time"
)

func AuthedUser(c *gin.Context) string {
	val, _ := c.Get("username")
	username := val.(string)
	return username
}
func DefaultFields(c *gin.Context) any {
	val, _ := c.Get("username")
	username := val.(string)
	fullName := "/SMS_APP/ADMIN/" + username
	//defaultFields :=

	return model.Default{
		CreatedBy:   fullName,
		UpdatedBy:   fullName,
		CreatedDate: time.Now().Local(),
		UpdatedDate: time.Now().Local(),
	}
}
func GetSession() gin.HandlerFunc {
	return func(context *gin.Context) {
		//return context
	}
}
