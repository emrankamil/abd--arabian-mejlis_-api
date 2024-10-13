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


func NewImageUploadRouter(env *bootstrap.Env, timeout time.Duration,db mongo.Database, group *gin.RouterGroup){
	productRepo := repository.NewProductRepository(db, domain.ProductsCollection, nil)
	productUsecase := usecase.NewProductUseCase(productRepo, timeout)
	uploadController := controller.NewUploadController(productUsecase, env)

	group.POST("/upload-image", uploadController.HandleUpload)
}