package mongodb

import (
	"context"
	"gohexdemo/internal/core/domains"

	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	db *mongo.Database
}

func NewOrderRepository(db *mongo.Database) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) SaveOrder(order *domains.Order) error {
	collection := r.db.Collection("orders")
	_, err := collection.InsertOne(context.Background(), order)
	return err
}
