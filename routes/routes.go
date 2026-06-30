package routes

import (
	"github.com/HRitsFadhila/golang-api-wallet/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	router := gin.Default()

	router.POST("/auth/register", controllers.Register)

	return  router
}