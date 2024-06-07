package Controllers

import (
	"github.com/gin-gonic/gin"
	"payment-microservice/Services"
	"payment-microservice/Utils"
)

type PaymentController struct {
	paymentService *Services.PaymentService
}

func NewPaymentController(paymentService *Services.PaymentService) *PaymentController {
	return &PaymentController{
		paymentService: paymentService,
	}
}

func (paymentController *PaymentController) CreatePayment(c *gin.Context) {
	var request Utils.PaymentCreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		Utils.RespondWithError(c, 400, "Invalid request")
		return
	}
	err := paymentController.paymentService.RegisterGateway(request)
	if err != nil {
		Utils.RespondWithError(c, 500, err.Error())
		return
	}
	response, err := paymentController.paymentService.CreatePayment(request)
	if err != nil {
		Utils.RespondWithError(c, 500, err.Error())
		return
	}
	Utils.RespondWithJSON(c, 200, response)

}

func (paymentController *PaymentController) GetPayment(c *gin.Context) {

}
