package route

import (
	"abduselam-arabianmejlis/bootstrap"
	"abduselam-arabianmejlis/delivery/controller"
	"abduselam-arabianmejlis/domain"
	"abduselam-arabianmejlis/infrastructure"
	"abduselam-arabianmejlis/mongo"
	"abduselam-arabianmejlis/repository"
	"abduselam-arabianmejlis/usecase"
	"time"

	"github.com/gin-gonic/gin"
)


func NewFogetPWRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup){
	ur := repository.NewUserRepository(db, domain.UserCollection)
	emailService := infrastructure.NewEmailService(env.SmtpServer, env.Mail, env.MailPassword)

	fpu := usecase.NewForgetPWUsecase(ur, timeout, *emailService)
	uu := usecase.NewUserUsecase(ur, timeout)

	fpc := controller.ForgetPWController{
		Userusecase: uu,
		ForgetPWUsecase: fpu,
		Env: env,
	}

	group.POST("/forget-password", fpc.ForgetPW)
	group.POST("/recover-password", fpc.ResetPW)
}