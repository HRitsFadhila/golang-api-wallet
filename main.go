package main

import (
	"github.com/HRitsFadhila/golang-api-wallet/config"
	"github.com/HRitsFadhila/golang-api-wallet/database"
	"github.com/HRitsFadhila/golang-api-wallet/routes"
)

func main(){
	config.LoadEnv()
	database.InitDB()
	
	r := routes.SetupRouter()

	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}