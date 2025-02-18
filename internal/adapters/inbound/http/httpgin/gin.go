package httpgin

import (
	"gohexdemo/internal/core/services"
	"net/http"
)

func StartServer(orderService *services.OrderService) {
	router := NewRouter(orderService)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
