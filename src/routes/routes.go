package routes

import (
	controller "app/src/controllers"
	"app/src/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	route.Use(middleware.Authenticate())
	user := route.Group("/user")
	{
		user.POST("/getUsers", controller.GetUsers())
		user.GET("/getUsers/:id", controller.GetUser())
		user.PUT("/getUsers/:id", controller.UpdateUser())
		user.PUT("/getUsers/create", controller.CreateUser())

		create := user.Group("/create")
		{
			step := create.Group("/step")
			{
				step.POST("/1", controller.StepOne())
				step.POST("/2", controller.StepTwo())
				step.POST("/3", controller.StepThree())
				step.POST("/3.1", controller.StepThreeOne())
			}
		}
	}
	settings := route.Group("/settings")
	{
		bannedWords := settings.Group("/bannedWords")
		{
			bannedWords.GET("/", controller.GetBannedWords())
			bannedWords.PUT("/", controller.PutBannedWords())
			bannedWords.DELETE("/", controller.DeleteBannedWords())
		}
		sms := settings.Group("/sms")
		{
			sms.GET("/limits", controller.GetSMSDailyLimit())
			sms.PUT("/limits", controller.UpdateSMSDailyLimit())
		}

	}
	getHardwareInfo := route.Group("/getHardware")
	{
		getHardwareInfo.POST("/MemoryInfo", controller.MemoryInfo())
		getHardwareInfo.POST("/CPUInfo", controller.CPUInfo())
		getHardwareInfo.POST("/HostInfo", controller.HostInfo())
		getHardwareInfo.POST("/InterfaceInfo", controller.InterfaceInfo())
		getHardwareInfo.POST("/DiskInfo", controller.DiskInfo())
	}
}
