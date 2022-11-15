package controller

import (
	"app/src/database"
	model "app/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CensoredWords struct {
	MN string `json:"mn"`
}

func GetBannedWords() gin.HandlerFunc {
	return func(context *gin.Context) {
		var BannedWords []CensoredWords
		var count int64
		database.DB.Find(&BannedWords).Count(&count)
		context.JSON(http.StatusOK, gin.H{
			"words": BannedWords,
			"count": count,
		})
	}
}
func PutBannedWords() gin.HandlerFunc {
	return func(context *gin.Context) {
		newWord := map[string]interface{}{}
		if err := context.BindJSON(&newWord); err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if err := database.DB.Table("censored_words").Create(newWord); err.Error != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error,
			})
			return
		}
		context.JSON(http.StatusOK, newWord)
	}
}
func DeleteBannedWords() gin.HandlerFunc {
	return func(context *gin.Context) {
		delWord := model.CensoredWords{}
		if err := context.BindJSON(&delWord); err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if err := database.DB.Model(&model.CensoredWords{}).Where("mn = ?", delWord.MN).Delete(delWord); err.Error != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error,
			})
			return
		}
		context.JSON(http.StatusOK, delWord)
	}
}

func GetSMSDailyLimit() gin.HandlerFunc {
	return func(context *gin.Context) {
		limits := map[string]interface{}{}
		database.DB.Table("sms_daily_limitation").Find(&limits)
		context.JSON(http.StatusOK, limits)
	}
}
func UpdateSMSDailyLimit() gin.HandlerFunc {
	return func(context *gin.Context) {
		var updateLimit map[string]interface{}
		if err := context.BindJSON(&updateLimit); err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if errDB := database.DB.Model(&model.SmsDailyLimitation{}).Where("status = ?", true).Updates(&updateLimit); errDB.Error != nil {
			context.JSON(http.StatusBadRequest, errDB.Error)
			return
		}
		context.JSON(http.StatusOK, updateLimit)
	}
}
