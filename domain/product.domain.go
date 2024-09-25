package domain

import (
	"context"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	ProductsCollection = "products"
	LikesCollection = "likes"
	ColorOptionsCollection = "color_options"
	ImageUploadFolder = "uploads"
	ImageQuality = 80  // The quality of the image to be uploaded out of 100

)

type Product struct {
	ID           	primitive.ObjectID  `json:"id" bson:"_id"`
	Title    	 	string              `json:"title" bson:"title" required:"true"`
	Title_AM     	string   			 `json:"title_am" bson:"title_am" required:"true"`
	Images       	[]string            `json:"images" bson:"images"`
	Description  	string              `json:"description" bson:"description" required:"true"`
	Description_AM  string				`json:"description_am" bson:"description_am" required:"true"`
	Category     	string    			 `json:"category" bson:"category" required:"true"`
	Features     	[]string            `json:"features" bson:"features"`
	Features_AM     []string            `json:"features_am" bson:"features_am"`
	Tags      	 	[]string            `json:"tags" bson:"tags"`
	IsAvailable  	bool      			 `json:"is_available" bson:"is_available"`
	Views        	int       			 `json:"views" bson:"views"`
	ColorOptions 	map[string]string `json:"color_options" bson:"color_options"`
	CreatedAt 	 	time.Time           `json:"created_at" bson:"created_at"`
	UpdatedAt 	 	time.Time           `json:"updated_at" bson:"updated_at"`
}

type Like struct {
	ID      	primitive.ObjectID `json:"id" bson:"_id"`
	UserID  	primitive.ObjectID `json:"user_id" bson:"user_id"`
	ProductID   primitive.ObjectID `json:"product_id" bson:"product_id"`
	IsLiked 	bool               `json:"is_liked" bson:"is_liked"`
}

type Pagination struct {
	Page     int `form:"page" json:"page"`
	PageSize int `form:"page_size" json:"page_size"`
}

type GetProductsRequest struct{
	Pagination  	`form:"pagination" bson:"pagination"`
	Category string	`form:"category" bson:"category"`
	Tag string		`form:"tag" bson:"tag"`
}

type ProductUseCase interface {
	CreateProduct(c context.Context, product *Product) (Product, error)
	GetProductByID(c context.Context, id string) (*Product,bool, error)
	GetProducts(c context.Context, pagination *Pagination, filter interface{}) ([]*Product, error)
	UpdateProduct(c context.Context, product *Product, product_id string) error
	DeleteProduct(c context.Context, id string) error
	UploadProductImages(ctx context.Context, files map[string]io.Reader, serverAdress string) ([]string, error)
	SearchProducts(ctx context.Context, query string) ([]*Product, error)
	LikeProduct(c context.Context, productID string, userID string) error
	UnlikeProduct(c context.Context, productID string, userID string) error
}

type ProductRepository interface {
	CreateProduct(c context.Context, Product *Product) (Product, error)
	GetProductByID(c context.Context, id string) (*Product, bool, error)
	GetProducts(c context.Context, pagination *Pagination, filter interface{}) ([]*Product, error)
	UpdateProduct(c context.Context, Product *Product) error
	DeleteProduct(c context.Context, id string) error
	SearchProducts(ctx context.Context, query string) ([]*Product, error)
	LikeProduct(c context.Context, ProductID string, userID string) error
	UnlikeProduct(c context.Context, ProductID string, userID string) error
}

type LikeRepository interface {
	LikeProduct(c context.Context, ProductID, userID primitive.ObjectID) error
	UnLikeProduct(c context.Context, ProductID, userID primitive.ObjectID) error
	DeleteLike(c context.Context, ProductID, userID primitive.ObjectID) error
	GetLike(ctx context.Context, userID, ProductID primitive.ObjectID) (*Like, error)
}

type LikeUsecase interface {
	LikeProduct(c context.Context, ProductID, userID primitive.ObjectID) error
	UnLikeProduct(c context.Context, ProductID, userID primitive.ObjectID) error
	DeleteLike(c context.Context, ProductID, userID primitive.ObjectID) error
	GetLike(ctx context.Context, userID, ProductID primitive.ObjectID) (*Like, error)
}