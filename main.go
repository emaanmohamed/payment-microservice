package main

import (
	"github.com/gin-gonic/gin"
	"payment-microservice/Config"
	"payment-microservice/Initializers"
	Routes "payment-microservice/Routes/Payment"
)

func main() {
	Initializers.ConnectToDB()
	server := gin.Default()
	routes := server.Group("/api/v1")
	Routes.SetUpPaymentRoutes(routes)
	server.Run(":" + Config.InitConfig().Port)

}
