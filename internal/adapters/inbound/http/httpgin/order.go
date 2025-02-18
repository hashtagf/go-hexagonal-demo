package httpgin

import (
	"gohexdemo/internal/core/domains"
	"gohexdemo/internal/core/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService *services.OrderService
}

func NewOrderController(orderService *services.OrderService) *OrderController {
	return &OrderController{orderService: orderService}
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	order := &domains.Order{}
	if err := ctx.ShouldBindJSON(order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.orderService.CreateOrder(order); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, order)
}
