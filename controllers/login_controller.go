package controllers

import (
	"net/http"

	"github.com/HRitsFadhila/golang-api-wallet/database"
	"github.com/HRitsFadhila/golang-api-wallet/helpers"
	"github.com/HRitsFadhila/golang-api-wallet/models"
	"github.com/HRitsFadhila/golang-api-wallet/structs"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context){
	var req = structs.UserLoginRequest{}
	var user = models.User{}

	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success : false,
			Message: "Validation Errors",
			Errors: helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil{
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "Email Not Found",
			Errors: helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil{
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "Invalid Password",
			Errors: helpers.TranslateErrorMessage(err),
		})
		return
	}

	token := helpers.GenerateToken(user.Email)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Login Success",
		Data: structs.UserResponse{
			Id: user.Id,
			Name: user.Name,
			Email: user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
			Token: &token,
		},
	})
}