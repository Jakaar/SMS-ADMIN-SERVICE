package controller

import (
	"app/src/database"
	helper "app/src/helpers"
	model "app/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	Username       string    `json:"username"`
	Firstname      string    `json:"firstname"`
	Lastname       string    `json:"lastname"`
	Avatar         string    `json:"avatar_url"`
	RegistryNumber string    `json:"registry_number"`
	UserType       string    `json:"user_type"`
	Phone          string    `json:"phone"`
	Email          string    `json:"email"`
	IsActive       bool      `json:"is_active"`
}

func GetUsers() gin.HandlerFunc {
	return func(context *gin.Context) {
		var users []User
		var count int64
		database.DB.Find(&users).Count(&count)
		context.JSON(http.StatusOK, gin.H{
			"count": count,
			"rows":  users,
		})
	}
}
func UpdateUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		updateUser := map[string]interface{}{
			"id": id,
		}
		if err := context.BindJSON(&updateUser); err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if err2 := database.DB.Model(&model.User{}).Where("id = ?", id).Updates(&updateUser).Error; err2 != nil {
			context.JSON(http.StatusBadRequest, err2.Error())
			return
		}
		context.JSON(http.StatusOK, updateUser)

	}
}
func GetUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		ID := context.Param("id")
		result := &model.User{}
		if err := database.DB.
			Where("id = ?", ID).
			First(&result).Error; err != nil {
			context.JSON(http.StatusOK, gin.H{
				"data": err.Error(),
			})
			return
		}
		context.JSON(http.StatusOK, result)
	}
}
func CreateUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		type NewUser struct {
			ID             uuid.UUID `json:"id"`
			WalletId       uuid.UUID `json:"wallet_id"`
			AvatarUrl      string    `json:"avatar_url"`
			Username       string    `json:"username"`
			Firstname      string    `json:"firstname" binding:"required"`
			Lastname       string    `json:"lastname"`
			RegistryNumber string    `json:"registry_number"`
			UserType       string    `json:"user_type"`
			AreaOfActivity string    `json:"area_of_activity"`
			Province       string    `json:"province"`
			District       string    `json:"district"`
			Committee      string    `json:"committee"`
			Address        string    `json:"address"`
			Email          string    `json:"email"`
			IsActive       bool      `json:"is_active"`
			IsVat          bool      `json:"is_vat"`
			Phone          int       `json:"phone"`
			Password       string    `json:"password"`
			model.Default
		}
		newWallet := model.Wallet{
			WalletId:           uuid.New(),
			CurrentBalance:     0,
			PostpaidLimitation: 0,
			PaymentSpending:    0,
		}
		newUser := NewUser{
			WalletId: newWallet.WalletId,
			IsActive: true,
		}
		if err := context.BindJSON(&newUser); err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}
		newUser.Password = newUser.RegistryNumber
		password, errP := helper.HashPassword(newUser.Password)

		if errP != nil {
			context.JSON(http.StatusBadRequest, errP.Error())
			return
		}
		newUser.Password = password
		newUser.ID = uuid.New()
		if err := database.DB.Create(&newWallet); err.Error != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error,
			})
			return
		}
		if err2 := database.DB.Table("user").Create(&newUser); err2.Error != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err2.Error,
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"new_user":   newUser,
			"new_wallet": newWallet,
		})
	}
}
func StepOne() gin.HandlerFunc {
	return func(context *gin.Context) {
		//session := sessions.Default(context)
		type NewUser struct {
			ID             uuid.UUID `json:"id"`
			WalletId       uuid.UUID `json:"wallet_id"`
			AvatarUrl      string    `json:"avatar_url"`
			Username       string    `json:"username"`
			Firstname      string    `json:"firstname" binding:"required"`
			Lastname       string    `json:"lastname"`
			RegistryNumber string    `json:"registry_number"`
			UserType       string    `json:"user_type"`
			AreaOfActivity string    `json:"area_of_activity"`
			Province       string    `json:"province"`
			District       string    `json:"district"`
			Committee      string    `json:"committee"`
			Address        string    `json:"address"`
			Email          string    `json:"email"`
			IsActive       bool      `json:"is_active"`
			IsVat          bool      `json:"is_vat"`
			Phone          int       `json:"phone"`
			Password       string    `json:"password"`
			model.Default
		}
		newWallet := model.Wallet{
			WalletId:           uuid.New(),
			CurrentBalance:     0,
			PostpaidLimitation: 0,
			PaymentSpending:    0,
		}
		newUser := NewUser{
			WalletId: newWallet.WalletId,
		}
		if err := context.BindJSON(&newUser); err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}
		password, errP := helper.HashPassword(newUser.Password)
		if errP != nil {
			context.JSON(http.StatusBadRequest, errP.Error())
			return
		}
		newUser.Password = password
		newUser.ID = uuid.New()
		if err := database.DB.Create(&newWallet); err.Error != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error,
			})
			return
		}
		if err2 := database.DB.Table("user").Create(&newUser); err2.Error != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err2.Error,
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"new_user":   newUser,
			"new_wallet": newWallet,
		})
	}
}
func StepTwo() gin.HandlerFunc {
	return func(context *gin.Context) {
		var price map[string]interface{}
		if err := context.BindJSON(&price); err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if err1 := database.DB.Table("price").Create(&price); err1.Error != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err1.Error,
			})
			return
		}
		context.JSON(http.StatusOK, price)
	}
}
func StepThree() gin.HandlerFunc {
	return func(context *gin.Context) {
		var merchant map[string]interface{}
		ID := uuid.New()
		if err := context.BindJSON(&merchant); err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}
		userID := merchant["user_id"].(string)

		create := model.Merchant{
			SelfID:       ID,
			Name:         merchant["name"].(string),
			UserID:       userID,
			MerchantCode: merchant["merchant_code"].(string),
		}
		if err1 := database.DB.Table("merchant").Create(&create); err1.Error != nil {
			context.JSON(http.StatusOK, err1.Error)
		}
		context.JSON(http.StatusOK, create)
	}
}
func StepThreeOne() gin.HandlerFunc {
	return func(context *gin.Context) {
		invoice := map[string]interface{}{}
		if err := context.BindJSON(&invoice); err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if err1 := database.DB.Table("invoice_template").Create(&invoice); err1.Error != nil {
			context.JSON(http.StatusOK, err1.Error)
		}
		context.JSON(http.StatusOK, invoice)
	}
}
