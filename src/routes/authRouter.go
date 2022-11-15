package routes

import (
	controller "app/src/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(route *gin.Engine) {
	route.POST("/login", controller.Login())
	//incoming.GET("/test", controller.Test())
}
