package ports

import (
	"gohexdemo/internal/core/domains"
)

// inbound port (primary port)
type InboundOrderPort interface {
	CreateOrder(order *domains.Order) error
}

// outbound port (secondary port)
type OrderRepository interface {
	SaveOrder(order *domains.Order) error
}
