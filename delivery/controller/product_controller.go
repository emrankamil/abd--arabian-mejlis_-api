package controller

import (
	"abduselam-arabianmejlis/domain"
	"abduselam-arabianmejlis/redis"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductController struct {
	ProductUseCase domain.ProductUseCase
	RedisClient    redis.Client
}

func NewProductController(pu domain.ProductUseCase, redisClient redis.Client) *ProductController {
	return &ProductController{
		ProductUseCase: pu,
		RedisClient:    redisClient,
	}
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var product domain.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
		return
	}
	product.ID = primitive.NewObjectID()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	product.IsAvailable = true
	product.Views = 13
	product.Likes = 8

	createdProduct, err := pc.ProductUseCase.CreateProduct(c, &product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, domain.SuccessResponse{Success: true, Data: createdProduct})
}

func (pc *ProductController) GetProductByID(c *gin.Context) {
	id := c.Param("id")

	product, fromCache, err := pc.ProductUseCase.GetProductByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	if product == nil {
		c.JSON(http.StatusNotFound, domain.SuccessResponse{Success: false, Data: nil})
		return
	}

	if fromCache {
		fmt.Println("from cache")
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Success: true, Data: product})

}

func (pc *ProductController) GetProducts(c *gin.Context) {
	var req domain.GetProductsRequest

	if err := c.BindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Error: "Invalid Request"})
		return
	}
	req.Page = max(req.Page, 1)
	req.PageSize = max(req.PageSize, 10)

	filter := bson.M{}
	if req.Category != "" {
		filter["category"] = req.Category
	}
	if req.Tag != "" {
		filter["tags"] = bson.M{"$in": []string{req.Tag}}
	}

	fmt.Println(req.Pagination, filter)
	products, totalCount, err := pc.ProductUseCase.GetProducts(c, &req.Pagination, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": products, "total_count": totalCount})

}

func (pc *ProductController) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product domain.Product
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Error: "Invalid Request"})
		return
	}
	err = pc.ProductUseCase.UpdateProduct(c, &product, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product Updated Successfully"})
}

func (pc *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	product, _, err := pc.ProductUseCase.GetProductByID(c, id)
	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product Not Found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	err = pc.ProductUseCase.DeleteProduct(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product Deleted Successfully"})
}

func (pc *ProductController) LikeProduct(c *gin.Context) {
	// productID := c.Param("product_id")
	// userID := c.Param("user_id")
	// err := pc.ProductUseCase.LikeProduct(c, productID, userID)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{"message": "Product Liked Successfully"})
	id := c.Param("id")
	product, _, err := pc.ProductUseCase.GetProductByID(c, id)
	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product Not Found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	product.Likes++
	err = pc.ProductUseCase.UpdateProduct(c, product, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product Liked Successfully", "likes": product.Likes, "success": true})

}

func (pc *ProductController) UnlikeProduct(c *gin.Context) {
	// productID := c.Param("product_id")
	// userID := c.Param("user_id")

	// err := pc.ProductUseCase.UnlikeProduct(c, productID, userID)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"message": "Product unliked successfully"})
	id := c.Param("id")
	product, _, err := pc.ProductUseCase.GetProductByID(c, id)
	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product Not Found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	product.Likes--
	err = pc.ProductUseCase.UpdateProduct(c, product, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product UnLiked Successfully", "likes": product.Likes, "success": true})
}

func (pc *ProductController) SearchProducts(c *gin.Context) {
	query := c.Query("q")
	products, err := pc.ProductUseCase.SearchProducts(c, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	if products == nil {
		c.JSON(http.StatusOK, gin.H{"sucess": false, "data": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": products})
}
