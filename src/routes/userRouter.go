package routes

import (
	"app/src/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine) {
	route.Use(middleware.Authenticate())
	//route.POST("/user/:id", controller.GetUsers())
}

//func AddUserRoutes(rg *gin.RouterGroup) {
//	users := rg.Group("/users")
//
//	users.POST("/", func(c *gin.Context) {
//		c.JSON(http.StatusOK, "users")
//	})
//	users.GET("/comments", func(c *gin.Context) {
//		c.JSON(http.StatusOK, "users comments")
//	})
//	users.GET("/pictures", func(c *gin.Context) {
//		c.JSON(http.StatusOK, "users pictures")
//	})
//}
