package controllers

import (
	"net/http"

	"github.com/HRitsFadhila/golang-api-wallet/database"
	"github.com/HRitsFadhila/golang-api-wallet/helpers"
	"github.com/HRitsFadhila/golang-api-wallet/models"
	"github.com/HRitsFadhila/golang-api-wallet/structs"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context){
	var req = structs.UserCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validasi Errors",
			Errors: helpers.TranslateErrorMessage(err),
		})
		return
	}
	
	tx := database.DB.Begin()
	defer func()  {
		if r := recover(); r != nil{
			tx.Rollback()
			panic(r)
		}
	}()

	user := models.User{
		Name: req.Name,
		Email: req.Email,
		Password: helpers.HashPassword(req.Password),
	}
	if err := tx.Create(&user).Error; err != nil{
		tx.Rollback()
			if helpers.IsDuplicateEntryError(err){
				c.JSON(http.StatusConflict, structs.ErrorResponse{
					Success: false,
					Message: "Duplicate entry error",
					Errors: helpers.TranslateErrorMessage(err),
				})
			} else {
				c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
					Success: false,
					Message: "Failed to create user",
					Errors: helpers.TranslateErrorMessage(err),
				})
			}
			return
	}

	account := models.Account{
		UserId: user.Id,
		AccountNumber: helpers.GenerateAccountNumber(),
		Balance: 0,
		Status: "active",
	}

	if err := tx.Create(&account).Error; err != nil{
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed create account",
			Errors: helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := tx.Commit().Error; err != nil{
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed save to database",
			Errors: helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "User created successfully",
		Data: gin.H{ 
			"user":structs.UserResponse{
				Id: user.Id,
				Name: user.Name,
				Email: user.Email,
				CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
			"account":structs.AccountResponse{
				AccountNumber: account.AccountNumber,
				Balance: helpers.FormatRupiah(account.Balance),
				Status: account.Status,
			},
		},
	})
}