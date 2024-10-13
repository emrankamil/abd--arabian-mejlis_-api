package usecase

import (
	"abduselam-arabianmejlis/bootstrap"
	"abduselam-arabianmejlis/domain"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type chatUsecase struct {
	chatRepo      domain.ChatRepository
	clientMgr *bootstrap.ClientManager
	timeout   time.Duration
}

// Read implements domain.ChatUsecase.
func (cu *chatUsecase) Read(receiverID string) {
	panic("unimplemented")
}

func NewChatUsecase(repo domain.ChatRepository, clientMgr *bootstrap.ClientManager, timeout time.Duration) domain.ChatUsecase {
	return &chatUsecase{
		chatRepo:      repo,
		clientMgr: clientMgr,
		timeout:   timeout,
	}
}

func (cu *chatUsecase) RegisterClient(client *domain.Client) {
	cu.clientMgr.Register <- client
}

func (cu *chatUsecase) UnregisterClient(client *domain.Client) {
	cu.clientMgr.Unregister <- client
}

func (cu *chatUsecase) ReadClientMessages(client *domain.Client, receiverID string) {
	defer func() {
		// Unregister the client and close the connection on function exit
		cu.clientMgr.Unregister <- client
		client.Conn.Close()
	}()

	for {
		messageType, p, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
		}

		message := domain.Message{
			ID:       primitive.NewObjectID(),
			Type:      messageType,
			Body:      string(p),
			SenderID:  client.UserID,
			RecipientID: receiverID,
			Timestamp: time.Now(),
		}

		cu.clientMgr.Broadcast <- message
		cu.chatRepo.CreateMessage(context.Background(), &message)
		log.Printf("Message Received: %+v\n", message)
	}
}

func (chu *chatUsecase) CreateMessage(c context.Context, message *domain.Message) error {
	ctx, cancel := context.WithTimeout(c, chu.timeout)
	defer cancel()
	return chu.chatRepo.CreateMessage(ctx, message)
}

func (chu *chatUsecase) DeleteMessage(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, chu.timeout)
	defer cancel()
	return chu.chatRepo.DeleteMessage(ctx, id)
}

func (chu *chatUsecase) GetMessagesByID(c context.Context, userID string, adminID string) ([]*domain.Message, error) {
	ctx, cancel := context.WithTimeout(c, chu.timeout)
	defer cancel()
	return chu.chatRepo.GetMessagesByID(ctx, userID, adminID)
}
