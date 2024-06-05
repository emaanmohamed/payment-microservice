package Utils

import "github.com/gin-gonic/gin"

type PaymentRequest struct {
	Gateway     string `json:"gateway" binding:"required"`
	Amount      int    `json:"amount" binding:"required"`
	Currency    string `json:"currency" binding:"required"`
	Description string `json:"description"`
}

type PaymentResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
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
