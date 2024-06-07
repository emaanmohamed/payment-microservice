package Utils

import "github.com/gin-gonic/gin"

type PaymentCreateRequest struct {
	Gateway     string `json:"gateway" binding:"required"`
	Amount      int    `json:"amount" binding:"required"`
	Currency    string `json:"currency" binding:"required"`
	Description string `json:"description"`
}

type PaymentRequest struct {
	ID      string `json:"ID" binding:"required"`
	Gateway string `json:"gateway" binding:"required"`
}

type PaymentResponse struct {
	Data    interface{} `json:"data"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
}

type PaymentStatusUpdateRequest struct {
	Status string `json:"status" binding:"required"`
}

func RespondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"error": message})
}

func RespondWithJSON(c *gin.Context, code int, payload interface{}) {
	c.JSON(code, payload)
}
