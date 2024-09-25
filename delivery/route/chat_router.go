package route

import (
	"abduselam-arabianmejlis/bootstrap"
	"abduselam-arabianmejlis/delivery/controller"
	"abduselam-arabianmejlis/domain"
	"abduselam-arabianmejlis/mongo"
	"abduselam-arabianmejlis/repository"
	"abduselam-arabianmejlis/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewChatRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, clientManager *bootstrap.ClientManager, group *gin.RouterGroup) {
	// Initialize the repository and usecase
	chatRepo := repository.NewChatRepository(db, domain.MessageCollection)
	chatUsecase := usecase.NewChatUsecase(chatRepo, clientManager, timeout)
	chatController := controller.NewChatController(chatUsecase)
	
	// WebSocket route
	group.GET("/ws", chatController.ServeWS)
}
