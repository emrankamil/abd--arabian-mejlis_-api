package usecase

import (
	"abduselam-arabianmejlis/domain"
	"abduselam-arabianmejlis/utils"
	"context"
	"io"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type productUseCase struct {
	likeRepo       domain.LikeRepository
	productRepo    domain.ProductRepository
	contextTimeout time.Duration
}

func NewProductUseCase(productRepo domain.ProductRepository, timeout time.Duration) domain.ProductUseCase {
	return &productUseCase{
		productRepo:    productRepo,
		contextTimeout: timeout,
	}
}

func (pu *productUseCase) CreateProduct(ctx context.Context, product *domain.Product) (domain.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()
	return pu.productRepo.CreateProduct(ctx, product)
}

func (pu *productUseCase) GetProductByID(ctx context.Context, id string) (*domain.Product, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()
	return pu.productRepo.GetProductByID(ctx, id)
}

func (pu *productUseCase) GetProducts(ctx context.Context, pagination *domain.Pagination, filter interface{}) ([]*domain.Product, int64, error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()
	return pu.productRepo.GetProducts(ctx, pagination, filter)
}

func (pu *productUseCase) UpdateProduct(ctx context.Context, product *domain.Product, productID string) error {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()
	// Ensure the product ID is set
	product.ID, _ = primitive.ObjectIDFromHex(productID)
	return pu.productRepo.UpdateProduct(ctx, product)
}

func (pu *productUseCase) DeleteProduct(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()
	return pu.productRepo.DeleteProduct(ctx, id)
}

func (pu *productUseCase) GetProductLikes(ctx context.Context, ProductID primitive.ObjectID) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()
	return pu.likeRepo.GetProductLikes(ctx, ProductID)
}

func (pu *productUseCase) LikeProduct(ctx context.Context, productID string, userID string) error {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)

	// convert the productID and userID to ObjectID
	productObjectID, _ := primitive.ObjectIDFromHex(productID)
	userObjectID, _ := primitive.ObjectIDFromHex(userID)

	defer cancel()
	return pu.likeRepo.LikeProduct(ctx, productObjectID, userObjectID)
}

func (pu *productUseCase) UnlikeProduct(ctx context.Context, productID string, userID string) error {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	productObjectID, _ := primitive.ObjectIDFromHex(productID)
	userObjectID, _ := primitive.ObjectIDFromHex(userID)
	defer cancel()
	return pu.likeRepo.UnLikeProduct(ctx, productObjectID, userObjectID)
}

func (pu *productUseCase) SearchProducts(ctx context.Context, keyword string) ([]*domain.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()
	return pu.productRepo.SearchProducts(ctx, keyword)
}

func (pu *productUseCase) UploadProductImages(ctx context.Context, files map[string]io.Reader, hostAddress string) ([]string, error) {
	var paths []string

	// Ensure the uploads directory exists
	errDir := utils.CreateFolder(domain.ImageUploadFolder)
	if errDir != nil {
		return nil, errDir
	}

	for _, file := range files {
		buffer, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}

		filename, err := utils.ImageProcessing(buffer, domain.ImageQuality, domain.ImageUploadFolder)
		if err != nil {
			return nil, err
		}

		filepath := "/" + domain.ImageUploadFolder + "/" + filename
		// Add the file path to the result
		paths = append(paths, filepath)
	}

	return paths, nil
}
