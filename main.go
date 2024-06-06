package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"payment-microservice/Config"
	Routes "payment-microservice/Routes/Payment"
)

func main() {
	config := Config.InitConfig()

	server := gin.Default()
	routes := server.Group("/api/v1")
	Routes.SetUpPaymentRoutes(routes)
	port := config.Port
	server.Run(fmt.Sprintf(":%s", port))

}
