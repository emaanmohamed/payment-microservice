package Routes

import (
	"github.com/gin-gonic/gin"
	"payment-microservice/Controllers"
	"payment-microservice/Middleware"
	"payment-microservice/Services"
)

var (
	PaymentService    = Services.NewPaymentService()
	PaymentController = Controllers.NewPaymentController(PaymentService)
)

func SetUpPaymentRoutes(router *gin.RouterGroup) {
	paymentGroup := router.Group("/Payment").Use(Middleware.AuthMiddleware())
	{
		paymentGroup.POST("/Create", PaymentController.CreatePayment)
		paymentGroup.POST("/Update", PaymentController.UpdatePaymentStatus)

		paymentGroup.POST("/Get", PaymentController.GetPayment).Use(Middleware.AdminMiddleware())
	}

}
