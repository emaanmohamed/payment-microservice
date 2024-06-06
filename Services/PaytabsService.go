package Services

import "payment-microservice/Utils"

type PaytabsService struct{}

func NewPaytabsService() *PaytabsService {
	return &PaytabsService{}
}

func (paytabsService *PaytabsService) GetPayment(id string) (Utils.PaymentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (paytabsService *PaytabsService) UpdatePaymentStatus(id string, request Utils.PaymentStatusUpdateRequest) (Utils.PaymentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (paytabsService *PaytabsService) CreatePayment(request Utils.PaymentRequest) (Utils.PaymentResponse, error) {
	// Call Paytabs API to create payment
	return Utils.PaymentResponse{
		//ID:     "123",
		Status: "success",
	}, nil

}
