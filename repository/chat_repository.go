package repository

import (
	"abduselam-arabianmejlis/domain"
	"abduselam-arabianmejlis/mongo"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type chatRepository struct {
	database   mongo.Database
	collection string
}

func NewChatRepository(db mongo.Database, collection string) domain.ChatRepository {
	return &chatRepository{
		database:   db,
		collection: collection,
	}
}
func (chr *chatRepository) CreateMessage(c context.Context, message *domain.Message) error {
	collection := chr.database.Collection(chr.collection)

	_, insertionErr := collection.InsertOne(c, message)
	return insertionErr
}

func (chr *chatRepository) DeleteMessage(c context.Context, id string) error {
	collection := chr.database.Collection(chr.collection)

	_, err := collection.DeleteOne(c, bson.M{"_id": id})
	return err
}

func (chr *chatRepository) GetMessagesByID(c context.Context, userID string, adminID string) ([]*domain.Message, error) {
	collection := chr.database.Collection(chr.collection)

	filter := bson.M{
        "$or": []bson.M{
            {"sender_id": adminID, "recipient_id": userID},
            {"sender_id": userID, "recipient_id": adminID},
        },
    }
	opts := options.Find().SetSort(bson.D{{Key: "timestamp", Value: -1}})
	
	cursor, err := collection.Find(c, filter, opts)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(c)

    var messages []*domain.Message
    if err = cursor.All(c, &messages); err != nil {
        return nil, err
    }
    return messages, nil
}


