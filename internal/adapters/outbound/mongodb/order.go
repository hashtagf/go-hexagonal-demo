package mongodb

import (
	"context"
	"gohexdemo/internal/core/domains"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ordersCollection = "orders"
	ordersDatabase   = "test_db"
)

type OrderRepository struct {
	db *mongo.Collection
}

func NewOrderRepository(db *mongo.Client) *OrderRepository {
	return &OrderRepository{db: db.Database(ordersDatabase).Collection(ordersCollection)}
}

func (r *OrderRepository) SaveOrder(order *domains.Order) error {
	_, err := r.db.InsertOne(context.Background(), order)
	return err
}
