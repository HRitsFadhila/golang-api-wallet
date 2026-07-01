package routes

import (
	"github.com/HRitsFadhila/golang-api-wallet/controllers"
	"github.com/HRitsFadhila/golang-api-wallet/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	router := gin.Default()

	router.POST("/auth/register", controllers.Register)
	router.POST("/auth/login", controllers.Login)

	router.GET("/ewallet/balance", middlewares.AuthMiddleware(), controllers.GetSaldo)

	return  router
}