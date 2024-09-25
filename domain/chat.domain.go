package domain

import (
	// websocket "abduselam-arabianmejlis/infrastructure/gorilla_websocket"
	"context"
	"time"

	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	MessageCollection = "messages"
)

type Client struct{
	ID				string
	Conn			*websocket.Conn
	UserID  		string
	IsAdmin 		bool
}

type Message struct{
    ID           primitive.ObjectID `bson:"_id" json:"_id"`
	Type		 int 				`json:"type" bson:"type"`
	Body		 string				`json:"body" bson:"body"`
	SenderID     string 			`json:"sender_id" bson:"sender_id"`
    RecipientID  string 			`json:"recipient_id" bson:"recipient_id"`
    Timestamp    time.Time  		`json:"timestamp" bson:"timestamp"`
}

// type ClientUsecase interface{
// }

type ChatUsecase interface{
	CreateMessage(c context.Context, message *Message) error
	GetMessagesByID(c context.Context, userID string, adminID string) ([]*Message, error)
	DeleteMessage(c context.Context, id string) error
	RegisterClient(client *Client)
	UnregisterClient(client *Client)
	ReadClientMessages(client *Client, receiverID string)
}

type ChatRepository interface{
	CreateMessage(c context.Context, message *Message) error
	GetMessagesByID(c context.Context, userID string, adminID string) ([]*Message, error)
	DeleteMessage(c context.Context, id string) error
}