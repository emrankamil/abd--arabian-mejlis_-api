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

func NewPromoteRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.UserCollection)
	pu := usecase.NewPromoteUsecase(ur, timeout)
	promoteUserController := controller.NewPromoteController(ur, pu, env)

	group.PUT("/promote-user/:id", promoteUserController.PromoteUser)
	group.PUT("/demote-user/:id", promoteUserController.DemoteUser)
}