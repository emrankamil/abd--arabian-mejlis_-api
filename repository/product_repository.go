package repository

import (
	"abduselam-arabianmejlis/domain"
	"abduselam-arabianmejlis/mongo"
	"abduselam-arabianmejlis/redis"
	"context"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type productRepository struct {
	database    mongo.Database
	collection  string
	RedisClient redis.Client
}

func NewProductRepository(db mongo.Database, collection string, redisClient redis.Client) domain.ProductRepository {
	return &productRepository{
		database:    db,
		collection:  collection,
		RedisClient: redisClient,
	}
}

func (r *productRepository) CreateProduct(c context.Context, product *domain.Product) (domain.Product, error) {
	// Insert the product into MongoDB
	collection := r.database.Collection(r.collection)
	_, err := collection.InsertOne(c, product)
	if err != nil {
		return *product, err
	}

	// Marshal the product to JSON for caching
	productData, err := json.Marshal(product)
	if err == nil {
		if r.RedisClient != nil {
			_ = r.RedisClient.Set(c, product.ID.Hex(), productData, 0)
		}

	}

	return *product, nil
}

func (r *productRepository) GetProductByID(c context.Context, id string) (*domain.Product, bool, error) {

	var product domain.Product
	if r.RedisClient != nil {
		cachedProduct, err := r.RedisClient.Get(c, id).Bytes()
		if err == nil {
			if err := json.Unmarshal(cachedProduct, &product); err == nil {
				return &product, true, nil
			}
		}
	}

	collection := r.database.Collection(r.collection)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, false, err
	}

	err = collection.FindOne(c, bson.M{"_id": objectID}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, false, nil // product not found
		}
		return nil, false, err
	}

	productData, err := json.Marshal(&product)
	if err == nil {
		if r.RedisClient != nil {
			_ = r.RedisClient.Set(c, id, productData, 0)
		}
	}

	// product.Views++
	// r.UpdateProduct(c, &product)
	return &product, false, nil
}

func (r *productRepository) GetProducts(c context.Context, pagination *domain.Pagination, filter interface{}) ([]*domain.Product, int64, error) {
	collection := r.database.Collection(r.collection)

	var products []*domain.Product

	skip := int64((pagination.Page - 1) * pagination.PageSize)
	limit := int64(pagination.PageSize)
	opts := &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	}
	cursor, err := collection.Find(c, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(c)

	for cursor.Next(c) {
		var product domain.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, 0, err
		}
		product.Views++
		r.UpdateProduct(c, &product)

		products = append(products, &product)
	}

	totalCount, err := collection.CountDocuments(c, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	return products, totalCount, nil
}

func (r *productRepository) UpdateProduct(c context.Context, product *domain.Product) error {
	collection := r.database.Collection(r.collection)

	_, err := collection.UpdateOne(
		c,
		bson.M{"_id": product.ID},
		bson.M{"$set": product},
	)
	if r.RedisClient != nil {
		r.RedisClient.Del(c, product.ID.Hex())
	}
	return err
}

func (r *productRepository) DeleteProduct(c context.Context, id string) error {
	collection := r.database.Collection(r.collection)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(c, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if r.RedisClient != nil {
		err = r.RedisClient.Del(c, id)
	}
	return err
}

func (r *productRepository) SearchProducts(c context.Context, keyword string) ([]*domain.Product, error) {
	collection := r.database.Collection(r.collection)

	var products []*domain.Product
	filter := bson.M{
		"$text": bson.M{
			"$search": keyword,
		},
	}

	cursor, err := collection.Find(c, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)

	for cursor.Next(c) {
		var product domain.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}

	return products, nil
}
