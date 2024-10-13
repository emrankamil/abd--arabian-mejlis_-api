package bootstrap

import (
	"fmt"
	"abduselam-arabianmejlis/domain"
)

// ClientManager handles the registration and broadcasting of messages
type ClientManager struct {
	Register   chan *domain.Client
	Unregister chan *domain.Client
	Clients    map[string]*domain.Client
	Broadcast  chan domain.Message
}

// NewClientManager creates a new instance of ClientManager
func NewClientManager() *ClientManager {
	return &ClientManager{
		Register:   make(chan *domain.Client),
		Unregister: make(chan *domain.Client),
		Clients:    make(map[string]*domain.Client),
		Broadcast:  make(chan domain.Message),
	}
}

func (manager *ClientManager) Start() {
	for {
		select {
		case client := <-manager.Register:
			manager.Clients[client.UserID] = client
			fmt.Println("Size of Connection Pool:", len(manager.Clients))

		case client := <-manager.Unregister:
			delete(manager.Clients, client.UserID)
			fmt.Println("Size of Connection Pool:", len(manager.Clients))

		case message := <-manager.Broadcast:
			fmt.Println("Sending message to recipient client in Manager")
			recipient, ok := manager.Clients[message.RecipientID]
			if ok {
				if err := recipient.Conn.WriteJSON(message); err != nil {
					fmt.Println("Error sending message:", err)
					return
				}
			}
		}
	}
}