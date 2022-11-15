package main

import (
	"app/src/database"
	model "app/src/models"
)

func init() {
	database.ConnectToDB()
}

func main() {
	//database.DB.AutoMigrate(&model.User{})
	database.DB.AutoMigrate(&model.CensoredWords{})
}
