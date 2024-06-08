package Utils

import "github.com/gin-gonic/gin"

type PaymentCreateRequest struct {
	Gateway     string  `json:"gateway" validate:"required"`
	Amount      float64 `json:"amount" validate:"required,gt=0"`
	Currency    string  `json:"currency" validate:"required,len=3"`
	Description string  `json:"description" validate:"required"`
	Status      string  `json:"status" validate:"required"`
}

type PaymentRequest struct {
	ID      string `json:"payment_id" binding:"required"`
	Gateway string `json:"gateway" binding:"required"`
}

type PaymentResponse struct {
	Data    interface{} `json:"data"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
}

type PaymentRegister struct {
	Gateway string `json:"gateway" binding:"required"`
}

type PaymentStatusUpdateRequest struct {
	ID      string `json:"payment_id" binding:"required"`
	Status  string `json:"status" binding:"required"`
	Gateway string `json:"gateway" binding:"required"`
}

func RespondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"error": message})
}

func RespondWithJSON(c *gin.Context, code int, payload interface{}) {
	c.JSON(code, payload)
}
