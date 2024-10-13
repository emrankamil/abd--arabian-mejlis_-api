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


func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup){
	ur := repository.NewUserRepository(db, domain.UserCollection)
	ru := usecase.NewRefreshTokenUsecase(ur, timeout)
	uu := usecase.NewUserUsecase(ur, timeout)
	rc := controller.NewRefreshTokenController(uu, ru, env)

	group.POST("/refresh_token", rc.RefreshTokenRequest)
}