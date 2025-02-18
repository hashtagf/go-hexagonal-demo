package elasticsearch

import (
	"context"
	"fmt"
	"gohexdemo/internal/core/domains"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

type OrderRepository struct {
	es *elasticsearch.Client
}

func NewOrderRepository(es *elasticsearch.Client) *OrderRepository {
	return &OrderRepository{es: es}
}

func (r *OrderRepository) SaveOrder(order *domains.Order) error {
	ctx := context.Background()

	orderJSON := order.ToJSON()

	// Create the document
	res, err := r.es.Index(
		"orders", // Index name
		strings.NewReader(orderJSON),
		r.es.Index.WithContext(ctx),
		r.es.Index.WithDocumentID(order.ID),
	)

	if err != nil {
		return fmt.Errorf("failed to index order: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error indexing order: %s", res.String())
	}

	return nil
}
