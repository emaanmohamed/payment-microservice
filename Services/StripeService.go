package Services

import "payment-microservice/Utils"

type StripeService struct{}

func NewStripeService() *StripeService {
	return &StripeService{}
}

func (stripeService *StripeService) GetPayment(id string) (Utils.PaymentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (stripeService *StripeService) UpdatePaymentStatus(id string, request Utils.PaymentStatusUpdateRequest) (Utils.PaymentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (stripeService *StripeService) CreatePayment(request Utils.PaymentRequest) (Utils.PaymentResponse, error) {
	// Call Stripe API to create payment
	return Utils.PaymentResponse{
		Status: "success",
	}, nil

}
