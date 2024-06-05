package Interfaces

import (
	"payment-microservice/Utils"
)

type Gateway interface {
	CreatePayment(request Utils.PaymentRequest) (Utils.PaymentResponse, error)
	GetPayment(id string) (Utils.PaymentResponse, error)
	UpdatePaymentStatus(id string, request Utils.PaymentStatusUpdateRequest) (Utils.PaymentResponse, error)
}
