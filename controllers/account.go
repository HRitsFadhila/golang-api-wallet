package controllers

import (
	"net/http"

	"github.com/HRitsFadhila/golang-api-wallet/database"
	"github.com/HRitsFadhila/golang-api-wallet/helpers"
	"github.com/HRitsFadhila/golang-api-wallet/models"
	"github.com/HRitsFadhila/golang-api-wallet/structs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetSaldo(c *gin.Context){
	userIdStr := c.MustGet("userId").(string)

	userId, err := uuid.Parse(userIdStr)
	if err != nil{
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "Ivalid user Id",
		})
		return
	}	
	var account models.Account

	if err := database.DB.Where("user_id = ?", userId).First(&account).Error; err != nil{
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Saldo not found",
			Errors: helpers.TranslateErrorMessage(err),
		})
	}
	
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Success get balance",
		Data: structs.BalanceResponse{
			Balance: helpers.FormatRupiah(account.Balance),
		},
	})
}