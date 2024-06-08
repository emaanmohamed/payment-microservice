package Services

import (
	"payment-microservice/Initializers"
	"payment-microservice/Utils"
	"payment-microservice/models"
)

type PaytabsService struct{}

func NewPaytabsService() *PaytabsService {
	return &PaytabsService{}
}

func (paytabsService *PaytabsService) GetPayment(id string) (Utils.PaymentResponse, error) {
	payment := models.Payment{}
	result := Initializers.DB.Where("id = ?", id).First(&payment)
	if result.Error != nil {
		return Utils.PaymentResponse{
			Status: "failed",
		}, result.Error
	}
	return Utils.PaymentResponse{
		Data:    payment,
		Message: "Payment retrieved successfully",
		Status:  "success",
	}, nil

}

func (paytabsService *PaytabsService) UpdatePaymentStatus(id string, request Utils.PaymentStatusUpdateRequest) (Utils.PaymentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (paytabsService *PaytabsService) CreatePayment(request Utils.PaymentCreateRequest) (Utils.PaymentResponse, error) {
	// Here I suggest to call the Paytabs API to create payment after successful payment I will add record to the database to verify the payment
	// I also expect that there is like unique transaction id that I can use to verify the payment and prevent duplicate payment
	payment := models.Payment{
		Gateway:     request.Gateway,
		Amount:      float64(request.Amount),
		Currency:    request.Currency,
		Description: request.Description,
		Status:      "success",
	}
	result := Initializers.DB.Create(&payment)
	if result.Error != nil {
		return Utils.PaymentResponse{
			Status: "failed",
		}, result.Error
	}
	return Utils.PaymentResponse{
		Data:    payment,
		Message: "Payment created successfully",
		Status:  "success",
	}, nil

}
