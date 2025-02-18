package services

import (
	"gohexdemo/internal/core/domains"
	"gohexdemo/internal/core/ports"
)

// business logic (Service)

type OrderService struct {
	orderRepository ports.OrderRepository
}

func NewOrderService(orderRepository ports.OrderRepository) *OrderService {
	return &OrderService{orderRepository: orderRepository}
}

func (s *OrderService) CreateOrder(order *domains.Order) error {
	if err := order.Validate(); err != nil {
		return err
	}
	return s.orderRepository.SaveOrder(order)
}
