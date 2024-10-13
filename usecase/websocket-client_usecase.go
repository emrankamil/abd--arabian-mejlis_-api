package usecase

import (
	"context"
	"log"
	"time"

	"abduselam-arabianmejlis/bootstrap"
	"abduselam-arabianmejlis/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type clientUsecase struct {
	client 			*domain.Client
	clientManager   *bootstrap.ClientManager
	chatRepository   domain.ChatRepository
}

// func NewClientUsecase(client *domain.Client, clientMgr *bootstrap.ClientManager, cr domain.ChatRepository) domain.ClientUsecase {
// 	return &clientUsecase{
// 		client: client,
// 		clientManager: clientMgr,
// 		chatRepository: cr,
// 	}
// }

func (cu *clientUsecase) Read(receiverID string){
	defer func() {
		// Unregister the client and close the connection on function exit
		cu.clientManager.Unregister <- cu.client
		cu.client.Conn.Close()
	}()

	for {
		// Read message from WebSocket connection
		messageType, p, err := cu.client.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
		}

		// Create a message struct with received data
		message := domain.Message{
			ID:       primitive.NewObjectID(),
			Type:      messageType,
			Body:      string(p),
			SenderID:  cu.client.UserID,
			RecipientID: receiverID,
			Timestamp: time.Now(),
		}

		cu.clientManager.Broadcast <- message
		cu.chatRepository.CreateMessage(context.Background(), &message)
		log.Printf("Message Received: %+v\n", message)
	}
}