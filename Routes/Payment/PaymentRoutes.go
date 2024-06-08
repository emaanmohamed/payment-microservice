package Routes

import (
	"github.com/gin-gonic/gin"
	"payment-microservice/Controllers"
	"payment-microservice/Services"
)

var (
	PaymentService    = Services.NewPaymentService()
	PaymentController = Controllers.NewPaymentController(PaymentService)
)

func SetUpPaymentRoutes(router *gin.RouterGroup) {
	paymentGroup := router.Group("/Payment")
	{
		paymentGroup.POST("/Create", PaymentController.CreatePayment)
		paymentGroup.POST("/Get", PaymentController.GetPayment)
		paymentGroup.POST("/Update", PaymentController.UpdatePaymentStatus)
	}

}
