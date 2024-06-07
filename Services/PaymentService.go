package Services

import (
	"errors"
	"log"
	"payment-microservice/Interfaces"
	"payment-microservice/Utils"
	"sync"
)

type PaymentService struct {
	gatewayRegistry map[string]Interfaces.Gateway
	mu              *sync.RWMutex
}

func NewPaymentService() *PaymentService {
	return &PaymentService{
		gatewayRegistry: make(map[string]Interfaces.Gateway),
		mu:              &sync.RWMutex{},
	}
}

func (paymentService *PaymentService) RegisterGateway(request Utils.PaymentCreateRequest) error {
	paymentService.mu.Lock()
	defer paymentService.mu.Unlock()

	if _, exists := paymentService.gatewayRegistry[request.Gateway]; exists {
		log.Printf("Gateway already registered: %s", request.Gateway)
		return nil
	}

	var gateway Interfaces.Gateway
	switch request.Gateway {
	case "stripe":
		gateway = NewStripeService()
	case "paytabs":
		gateway = NewPaytabsService()
	default:
		return errors.New("unsupported gateway type")
	}

	paymentService.gatewayRegistry[request.Gateway] = gateway
	log.Printf("Registered gateway: %s", request.Gateway)
	return nil
}

func (paymentService *PaymentService) GetPaymentService(gateway string) (Interfaces.Gateway, error) {
	paymentService.mu.Lock()
	defer paymentService.mu.Unlock()

	service, exists := paymentService.gatewayRegistry[gateway]
	if !exists {
		return nil, errors.New("gateway not found")
	}
	return service, nil

}

func (paymentService *PaymentService) CreatePayment(request Utils.PaymentCreateRequest) (Utils.PaymentResponse, error) {
	service, error := paymentService.GetPaymentService(request.Gateway)
	if error != nil {
		return Utils.PaymentResponse{}, error
	}
	return service.CreatePayment(request)

}

func (paymentService *PaymentService) GetPayment(id string, gateway string) (Utils.PaymentResponse, error) {
	service, error := paymentService.GetPaymentService(gateway)
	if error != nil {
		return Utils.PaymentResponse{}, error
	}
	return service.GetPayment(id)
}

func (paymentService *PaymentService) UpdatePaymentStatus(id string, request Utils.PaymentStatusUpdateRequest, gateway string) (Utils.PaymentResponse, error) {
	service, error := paymentService.GetPaymentService(gateway)
	if error != nil {
		return Utils.PaymentResponse{}, error
	}
	return service.UpdatePaymentStatus(id, request)
}
