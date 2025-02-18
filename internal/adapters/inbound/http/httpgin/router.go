package httpgin

import (
	"gohexdemo/internal/core/services"

	"github.com/gin-gonic/gin"
)

func NewRouter(orderService *services.OrderService) *gin.Engine {
	router := gin.Default()
	orderController := NewOrderController(orderService)
	router.POST("/orders", orderController.CreateOrder)
	return router
}
