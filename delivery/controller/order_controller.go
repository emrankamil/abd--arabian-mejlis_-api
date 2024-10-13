package controller

import (
	"abduselam-arabianmejlis/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderController struct {
	OrderUseCase domain.OrderUseCase
}

func NewOrderController(ou domain.OrderUseCase) *OrderController {
	return &OrderController{
		OrderUseCase: ou,
	}
}

func (oc *OrderController) CreateOrder(c *gin.Context) {
	var order domain.Order
	err := c.ShouldBindJSON(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
		return
	}
	order.ID = primitive.NewObjectID()
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	createdOrder, err := oc.OrderUseCase.CreateOrder(c, &order)

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, domain.SuccessResponse{Success: true, Data: createdOrder})
}

func (oc *OrderController) GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	order, err := oc.OrderUseCase.GetOrderByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	if order == nil {
		c.JSON(http.StatusNotFound, domain.SuccessResponse{Success: false, Data: nil})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Success: true, Data: order})
}

func (oc *OrderController) GetOrders(c *gin.Context) {
	productID := c.Query("product_id")
	email := c.Query("email")

	orders, err := oc.OrderUseCase.GetOrders(c, email, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Success: true, Data: orders})
}

func (oc *OrderController) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	err := oc.OrderUseCase.DeleteOrder(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Success: true, Data: nil})
}
