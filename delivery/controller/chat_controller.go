package controller

import (
	"abduselam-arabianmejlis/domain"
	websocket "abduselam-arabianmejlis/infrastructure/gorilla_websocket"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChatController struct {
	ChatUsecase domain.ChatUsecase
}

func NewChatController(cu domain.ChatUsecase) *ChatController {
	return &ChatController{
		ChatUsecase: cu,
	}
}

func (cc *ChatController) ServeWS(c *gin.Context) {
	receiverID := c.Query("receiver")
	userID := c.Query("user_id")
	userType := c.Query("user_type")
	// userID, exists := c.Get("user_id")
	// if !exists {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized: missing user_id"})
	// 	return
	// }
	
	// userType, exists := c.Get("user_type")
	// if !exists {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized: missing user_type"})
	// 	return
	// }

	isAdmin := userType == "ADMIN"
	
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := websocket.Upgrade(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	
	// Initialize the client
	client := &domain.Client{
		ID:      userID,
		Conn:    conn,
		UserID:  userID,
		IsAdmin: isAdmin,
	}
	// Register the client with the chat usecase
	cc.ChatUsecase.RegisterClient(client)

	// Start reading messages for the client
	cc.ChatUsecase.ReadClientMessages(client, receiverID)
}
