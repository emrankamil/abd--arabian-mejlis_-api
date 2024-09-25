package controller

import (
	"abduselam-arabianmejlis/domain"
	"abduselam-arabianmejlis/redis"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductController struct {
	ProductUseCase domain.ProductUseCase
	RedisClient redis.Client
}

func NewProductController(pu domain.ProductUseCase, redisClient redis.Client) *ProductController {
	return &ProductController{
		ProductUseCase: pu,
		RedisClient: redisClient,
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

	createdProduct, err := pc.ProductUseCase.CreateProduct(c, &product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, domain.SuccessResponse{Success: true, Message: "Product Created Successfully", Data: createdProduct})
}

func (pc *ProductController) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	
    product, fromCache, err := pc.ProductUseCase.GetProductByID(c, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
        return
    }
	if product == nil {
		c.JSON(http.StatusNotFound, domain.SuccessResponse{Success: false, Message: "Product Not Found", Data: nil})
		return
	}

    message := "Product Fetched Successfully"
    if fromCache {
        message = "Product Fetched From Cache Successfully"
    }

    c.JSON(http.StatusOK, domain.SuccessResponse{Success: true, Message: message, Data: product})

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
		filter["tags"] =  bson.M{"$in": []string{req.Tag}}
	}

	products, err := pc.ProductUseCase.GetProducts(c, &req.Pagination, filter)
    if err != nil {
        c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
        return
    }

    c.JSON(http.StatusOK, domain.SuccessResponse{Success: true, Message: "Products Fetched Successfully", Data: products})
	
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
	if product == nil{
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
	productID := c.Param("product_id")
	userID := c.Param("user_id")
	err := pc.ProductUseCase.LikeProduct(c, productID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product Liked Successfully"})
}

func (pc *ProductController) UnlikeProduct(c *gin.Context) {
	productID := c.Param("product_id")
	userID := c.Param("user_id")

	err := pc.ProductUseCase.UnlikeProduct(c, productID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product unliked successfully"})
}

func (pc *ProductController) SearchProducts(c *gin.Context) {
	query := c.Query("q")
	products, err := pc.ProductUseCase.SearchProducts(c, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	if products == nil {
		c.JSON(http.StatusOK, gin.H{"message": "No products found"})
		return
	}
	c.JSON(http.StatusOK, products)
}
