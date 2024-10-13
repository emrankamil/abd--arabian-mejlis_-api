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

func NewOrderRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	orderRepo := repository.NewOrderRepository(db, domain.OrdersCollection)
	orderUsecase := usecase.NewOrderUseCase(orderRepo, timeout)
	OrderController := controller.NewOrderController(orderUsecase)

	group.POST("/order", OrderController.CreateOrder)
	group.GET("/order/:id", OrderController.GetOrderByID)
	group.GET("/order", OrderController.GetOrders)
	group.DELETE("/order/:id", OrderController.DeleteOrder)

}
