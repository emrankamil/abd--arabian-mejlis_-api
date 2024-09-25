package route

import (
	"abduselam-arabianmejlis/bootstrap"
	"abduselam-arabianmejlis/delivery/controller"
	"abduselam-arabianmejlis/domain"
	"abduselam-arabianmejlis/mongo"
	"abduselam-arabianmejlis/redis"
	"abduselam-arabianmejlis/repository"
	"abduselam-arabianmejlis/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewProductRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup, redisClient redis.Client) {
	productRepo := repository.NewProductRepository(db, domain.ProductsCollection, redisClient)
	productUsecase := usecase.NewProductUseCase(productRepo, timeout)
	productController := controller.NewProductController(productUsecase, redisClient)

	group.POST("/products", productController.CreateProduct)
	group.GET("/products/:id", productController.GetProductByID)
	group.GET("/products", productController.GetProducts)
	group.PUT("/products/:id", productController.UpdateProduct)
	group.DELETE("/products/:id", productController.DeleteProduct)
	group.GET("/products/search", productController.SearchProducts)
	group.POST("/products/:id/like", productController.LikeProduct)
	group.POST("/products/:id/unlike", productController.UnlikeProduct)
}
