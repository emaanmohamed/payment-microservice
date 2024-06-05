package Services

import (
	"errors"
	"payment-microservice/Interfaces"
	"payment-microservice/Utils"
)

var gatewayRegistry = make(map[string]Interfaces.Gateway)

func RegisterGateway(name string, gateway Interfaces.Gateway) {
	gatewayRegistry[name] = gateway
}

func GetPaymentService(gateway string) (Interfaces.Gateway, error) {
	service, ok := gatewayRegistry[gateway]
	if !ok {
		return nil, errors.New("gateway not found")
	}
	return service, nil

}

type PaymentService struct{}

func (s *PaymentService) CreatePayment(request Utils.PaymentRequest) (Utils.PaymentResponse, error) {
	service, error := GetPaymentService(request.Gateway)

}
