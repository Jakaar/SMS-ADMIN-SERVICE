package controller

import (
	"app/src/database"
	helper "app/src/helpers"
	model "app/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login() gin.HandlerFunc {
	return func(context *gin.Context) {
		var data struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := context.ShouldBindJSON(&data); err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		admin := model.Admin{}
		if data.Username == "" {
			context.JSON(400, "error")
			return
		}
		database.DB.Where("username = ?", data.Username).Find(&admin)
		if admin.Status == false {
			context.JSON(http.StatusUnauthorized, "no credential")
			return
		}
		if check := helper.CheckPasswordHash(data.Password, admin.Password); check != true {
			context.JSON(http.StatusOK, "invalid username or password")
			return
		}
		token, refreshToken, err := helper.GenerateToken(admin.Username, admin.ID)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		helper.UpdateToken(token, refreshToken, admin.ID)
		context.JSON(http.StatusOK, gin.H{
			"access_token":  token,
			"refresh_token": refreshToken,
		})
	}
}
