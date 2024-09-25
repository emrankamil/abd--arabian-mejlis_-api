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


func NewLogoutRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup){
	ur := repository.NewUserRepository(db, domain.UserCollection)
	su := usecase.NewLoginUsecase(ur, timeout)
	uu := usecase.NewUserUsecase(ur, timeout)
	sc := controller.LoginController{
		UserUsecase: uu,
		LoginUsecase: su,
		Env: env,
	}
	
	group.GET("/logout", sc.Logout)
}