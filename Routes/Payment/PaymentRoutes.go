package Routes

import "github.com/gin-gonic/gin"

func SetUpPaymentRoutes(router *gin.RouterGroup) {
	paymentGroup := router.Group("/Payment")
	{
		paymentGroup.POST("/Pay", Pay)
	}

}
