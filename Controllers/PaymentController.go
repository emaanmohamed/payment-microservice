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

	gateway := Utils.PaymentRegister{Gateway: request.Gateway}

	err := paymentController.paymentService.RegisterGateway(gateway)
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
	var request Utils.PaymentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		Utils.RespondWithError(c, 400, "Invalid request format")
		return
	}

	gateway := Utils.PaymentRegister{Gateway: request.Gateway}

	err := paymentController.paymentService.RegisterGateway(gateway)
	if err != nil {
		Utils.RespondWithError(c, 500, err.Error())
		return
	}
	response, err := paymentController.paymentService.GetPayment(request.ID, request.Gateway)
	if err != nil {
		Utils.RespondWithError(c, 500, err.Error())
		return
	}
	Utils.RespondWithJSON(c, 200, response)

}

func (paymentController *PaymentController) UpdatePaymentStatus(c *gin.Context) {
	var request Utils.PaymentStatusUpdateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		Utils.RespondWithError(c, 400, "Invalid request format")
		return
	}

	gateway := Utils.PaymentRegister{Gateway: request.Gateway}

	err := paymentController.paymentService.RegisterGateway(gateway)
	if err != nil {
		Utils.RespondWithError(c, 500, err.Error())
		return
	}
	response, err := paymentController.paymentService.UpdatePaymentStatus(request.ID, request, request.Gateway)
	if err != nil {
		Utils.RespondWithError(c, 500, err.Error())
		return
	}
	Utils.RespondWithJSON(c, 200, response)

}

func (paymentController *PaymentController) RegisterGateway(request Utils.PaymentRegister) (Utils.PaymentRegister, error) {
	gateway := Utils.PaymentRegister{Gateway: request.Gateway}

	err := paymentController.paymentService.RegisterGateway(gateway)
	if err != nil {
		return gateway, err

	}
	return gateway, nil
}
