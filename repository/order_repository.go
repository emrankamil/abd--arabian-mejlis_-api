package repository

import (
	"abduselam-arabianmejlis/domain"
	"abduselam-arabianmejlis/mongo"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type orderRepository struct {
	database   mongo.Database
	collection string
}

func NewOrderRepository(db mongo.Database, collection string) domain.OrderRepository {
	return &orderRepository{
		database:   db,
		collection: collection,
	}
}

func (or *orderRepository) CreateOrder(c context.Context, order *domain.Order) (domain.Order, error) {
	collection := or.database.Collection(or.collection)
	_, err := collection.InsertOne(c, order)
	if err != nil {
		return *order, err
	}
	return *order, nil
}

func (or *orderRepository) DeleteOrder(c context.Context, id string) error {
	collection := or.database.Collection(or.collection)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(c, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	return nil
}

func (or *orderRepository) GetOrderByEmail(c context.Context, email string) ([]*domain.Order, error) {
	var orders []*domain.Order
	collection := or.database.Collection(or.collection)

	cursor, err := collection.Find(c, bson.M{"email": email})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	for cursor.Next(c) {
		var order domain.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}

		orders = append(orders, &order)
	}

	return orders, nil
}

func (or *orderRepository) GetOrderByProductID(c context.Context, productID string) ([]*domain.Order, error) {
	var orders []*domain.Order
	collection := or.database.Collection(or.collection)
	objectID, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		return nil, err
	}

	cursor, err := collection.Find(c, bson.M{"product_id": objectID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	for cursor.Next(c) {
		var order domain.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}

		orders = append(orders, &order)
	}

	return orders, nil
}

func (or *orderRepository) GetOrders(c context.Context, filter interface{}) ([]*domain.Order, error) {
	var orders []*domain.Order
	collection := or.database.Collection(or.collection)

	// Use the filter to find matching orders
	cursor, err := collection.Find(c, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	for cursor.Next(c) {
		var order domain.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}

		orders = append(orders, &order)
	}

	return orders, nil
}

func (or *orderRepository) GetOrderByID(c context.Context, id string) (*domain.Order, error) {
	var order domain.Order
	collection := or.database.Collection(or.collection)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = collection.FindOne(c, bson.M{"_id": objectID}).Decode(&order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // product not found
		}
		return nil, err
	}

	return &order, nil
}
