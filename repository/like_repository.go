package repository

import (
	"abduselam-arabianmejlis/domain"
	"abduselam-arabianmejlis/mongo"
	"abduselam-arabianmejlis/redis"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LikeRepository struct {
	database    mongo.Database
	collection  string
	RedisClient redis.Client
}

// DeleteLike implements domain.LikeRepository.
func (l *LikeRepository) DeleteLike(c context.Context, ProductID primitive.ObjectID, userID primitive.ObjectID) error {
	panic("unimplemented")
}

// GetLike implements domain.LikeRepository.
func (l *LikeRepository) GetProductLikes(ctx context.Context, ProductID primitive.ObjectID) (int64, error) {
	collection := l.database.Collection(domain.LikesCollection)

	likes, err := collection.CountDocuments(ctx, bson.M{"_id": ProductID})
	if err != nil {
		return 0, err
	}

	return likes, nil
}

// LikeProduct implements domain.LikeRepository.
func (l *LikeRepository) LikeProduct(c context.Context, ProductID primitive.ObjectID, userID primitive.ObjectID) error {
	collection := l.database.Collection(domain.LikesCollection)

	filter := bson.M{"_id": ProductID}
	update := bson.M{"$addToSet": bson.M{"likes": userID}}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

// UnLikeProduct implements domain.LikeRepository.
func (l *LikeRepository) UnLikeProduct(c context.Context, ProductID primitive.ObjectID, userID primitive.ObjectID) error {
	collection := l.database.Collection(domain.LikesCollection)

	filter := bson.M{"_id": ProductID}
	update := bson.M{"$pull": bson.M{"likes": userID}}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}

func NewLikeRepository(db mongo.Database, collection string, redisClient redis.Client) domain.LikeRepository {
	return &LikeRepository{
		database:    db,
		collection:  collection,
		RedisClient: redisClient,
	}
}