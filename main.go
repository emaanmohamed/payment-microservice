package main

import (
	"github.com/gin-gonic/gin"
	"payment-microservice/Config"
	"payment-microservice/Initializers"
	RoutesAuth "payment-microservice/Routes/Auth"
	RoutesPayment "payment-microservice/Routes/Payment"
)

func main() {
	Initializers.ConnectToDB()
	server := gin.Default()
	routes := server.Group("/api/v1")
	RoutesPayment.SetUpPaymentRoutes(routes)
	RoutesAuth.SetUpAuthRoutes(routes)
	server.Run(":" + Config.InitConfig().Port)

}
