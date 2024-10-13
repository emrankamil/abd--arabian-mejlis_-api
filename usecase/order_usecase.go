package usecase

import (
	"abduselam-arabianmejlis/domain"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type orderUsecase struct {
	orderRepo      domain.OrderRepository
	contextTimeout time.Duration
}

func NewOrderUseCase(orderRepo domain.OrderRepository, timeout time.Duration) domain.OrderUseCase {
	return &orderUsecase{
		orderRepo:      orderRepo,
		contextTimeout: timeout,
	}
}

func (ou *orderUsecase) CreateOrder(c context.Context, order *domain.Order) (domain.Order, error) {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	return ou.orderRepo.CreateOrder(ctx, order)
}

func (ou *orderUsecase) DeleteOrder(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	return ou.orderRepo.DeleteOrder(ctx, id)
}

func (ou *orderUsecase) GetOrderByEmail(c context.Context, email string) ([]*domain.Order, error) {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	return ou.orderRepo.GetOrderByEmail(ctx, email)
}

func (ou *orderUsecase) GetOrders(c context.Context, email, productID string) ([]*domain.Order, error) {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()

	var filter bson.M

	if productID != "" {
		objectID, err := primitive.ObjectIDFromHex(productID)
		if err != nil {
			return nil, err
		}
		filter = bson.M{"product_id": objectID}
	} else if email != "" {
		filter = bson.M{"email": email}
	} else {
		filter = bson.M{}
	}

	return ou.orderRepo.GetOrders(ctx, filter)
}

func (ou *orderUsecase) GetOrderByID(c context.Context, id string) (*domain.Order, error) {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	return ou.orderRepo.GetOrderByID(ctx, id)
}

func (ou *orderUsecase) GetOrderByProductID(c context.Context, productID string) ([]*domain.Order, error) {
	ctx, cancel := context.WithTimeout(c, ou.contextTimeout)
	defer cancel()
	return ou.orderRepo.GetOrderByProductID(ctx, productID)
}
