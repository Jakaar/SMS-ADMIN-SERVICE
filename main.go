package main

import (
	"app/src/database"
	"app/src/middleware"
	"app/src/routes"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"os"
)

const (
	VERSION = "0.1"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.JSONLoggerMiddleware())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())
	store := cookie.NewStore([]byte(os.Getenv("TOKEN_SECRET_KEY")))
	r.Use(sessions.Sessions("userSession", store))
	routes.AuthRoutes(r)
	routes.Routes(r)
	return r
}
func init() {
	err := database.ConnectToDB()
	if err != nil {
		return
	}
}
func main() {
	//helper.Log.Printf("Server v%s pid=%d started with processes: %d", VERSION, )
	r := setupRouter()
	r.Run(":" + os.Getenv("PORT"))
}
