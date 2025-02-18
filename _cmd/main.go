package main

import (
	"fmt"
	"gohexdemo"
	"gohexdemo/internal/adapters/outbound/mongodb"
	"gohexdemo/internal/core/services"
	"log"
)

func main() {
	fmt.Println("Hello Hexagonal Architecture")

	// connect to mongodb
	db, err := mongodb.Connect()
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := mongodb.NewOrderRepository(db.Database("orders"))
	orderService := services.NewOrderService(orderRepository)
	gohexdemo.StartServer(orderService)
}
