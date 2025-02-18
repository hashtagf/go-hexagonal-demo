package gohexdemo

import (
	httpgin "gohexdemo/internal/adapters/inbound/http/httpgin"
	"gohexdemo/internal/core/services"
)

func StartServer(orderService *services.OrderService) {
	httpgin.StartServer(orderService)
}
