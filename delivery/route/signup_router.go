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

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup){
	emailService := infrastructure.NewEmailService(env.SmtpServer, env.Mail, env.MailPassword)
	ur := repository.NewUserRepository(db, domain.UserCollection)
	su := usecase.NewSignupUsecase(ur, timeout, *emailService)
	uu := usecase.NewUserUsecase(ur, timeout)
	sc := controller.NewSignupController(uu, su, env)

	group.POST("/signup", sc.Signup)
	group.POST("/verify_email", sc.VerifyEmail)
}