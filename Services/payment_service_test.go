package Services

import (
	"github.com/stretchr/testify/assert"
	"payment-microservice/Utils"
	"testing"
)

type MockGateway struct{}

func (m *MockGateway) CreatePayment(request Utils.PaymentCreateRequest) (Utils.PaymentResponse, error) {
	return Utils.PaymentResponse{Data: "123", Status: "success", Message: "Payment created successfully"}, nil
}

func (m *MockGateway) GetPayment(id string) (Utils.PaymentResponse, error) {
	return Utils.PaymentResponse{Data: id, Status: "success", Message: "Payment retrieved successfully"}, nil
}

func (m *MockGateway) UpdatePaymentStatus(id string, request Utils.PaymentStatusUpdateRequest) (Utils.PaymentResponse, error) {
	return Utils.PaymentResponse{Data: id, Status: request.Status, Message: "Payment status updated successfully"}, nil
}

func TestRegisterGateway(t *testing.T) {
	paymentService := NewPaymentService()

	request := Utils.PaymentRegister{Gateway: "paytabs"}
	err := paymentService.RegisterGateway(request)
	assert.NoError(t, err)
	assert.Contains(t, paymentService.gatewayRegistry, "paytabs")

	err = paymentService.RegisterGateway(request)
	assert.NoError(t, err)
	assert.Contains(t, paymentService.gatewayRegistry, "paytabs")
}

func TestGetPaymentService(t *testing.T) {
	paymentService := NewPaymentService()
	paymentService.gatewayRegistry["paytabs"] = &MockGateway{}

	service, err := paymentService.GetPaymentService("paytabs")
	assert.NoError(t, err)
	assert.NotNil(t, service)

	service, err = paymentService.GetPaymentService("nonexistent")
	assert.Error(t, err)
	assert.Nil(t, service)
}

func TestCreatePayment(t *testing.T) {
	paymentService := NewPaymentService()
	paymentService.gatewayRegistry["paytabs"] = &MockGateway{}

	request := Utils.PaymentCreateRequest{Gateway: "paytabs"}
	response, err := paymentService.CreatePayment(request)
	assert.NoError(t, err)
	assert.Equal(t, "123", response.Data)
	assert.Equal(t, "success", response.Status)
	assert.Equal(t, "Payment created successfully", response.Message)

	request.Gateway = "nonexistent"
	response, err = paymentService.CreatePayment(request)
	assert.Error(t, err)
	assert.Nil(t, response.Data)
}

func TestGetPayment(t *testing.T) {
	paymentService := NewPaymentService()
	paymentService.gatewayRegistry["paytabs"] = &MockGateway{}

	response, err := paymentService.GetPayment("123", "paytabs")
	assert.NoError(t, err)
	assert.Equal(t, "123", response.Data)
	assert.Equal(t, "success", response.Status)
	assert.Equal(t, "Payment retrieved successfully", response.Message)

	response, err = paymentService.GetPayment("123", "nonexistent")
	assert.Error(t, err)
	assert.Nil(t, response.Data)
}

func TestUpdatePaymentStatus(t *testing.T) {
	paymentService := NewPaymentService()
	paymentService.gatewayRegistry["paytabs"] = &MockGateway{}

	request := Utils.PaymentStatusUpdateRequest{Status: "updated"}
	response, err := paymentService.UpdatePaymentStatus("123", request, "paytabs")
	assert.NoError(t, err)
	assert.Equal(t, "123", response.Data)
	assert.Equal(t, "updated", response.Status)
	assert.Equal(t, "Payment status updated successfully", response.Message)

	response, err = paymentService.UpdatePaymentStatus("123", request, "nonexistent")
	assert.Error(t, err)
	assert.Nil(t, response.Data)
}
