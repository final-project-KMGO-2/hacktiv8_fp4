package service

import (
	"context"
	"hacktiv8_fp_2/entity"
	"hacktiv8_fp_2/repository"

	"github.com/mashingan/smapping"
)

type ProductService interface {
	CreateNewProduct(ctx context.Context, productCrt entity.ProductCreate) (entity.Product, error)
	GetAllProducts(ctx context.Context) ([]entity.Product, error)
	UpdateProductById(ctx context.Context, productUpd entity.ProductUpdate, id uint64) (entity.Product, error)
	DeleteProductById(ctx context.Context, id uint64) error
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(pr repository.ProductRepository) ProductService {
	return &productService{productRepo: pr}
}

// TODO: Check category every create and update

func (pr *productService) CreateNewProduct(ctx context.Context, productCrt entity.ProductCreate) (entity.Product, error) {
	product := entity.Product{}
	err := smapping.FillStruct(&product, smapping.MapFields(&productCrt))
	if err != nil {
		return entity.Product{}, err
	}

	// check category di sini

	result, err := pr.productRepo.InsertNewProduct(ctx, product)
	if err != nil {
		return entity.Product{}, err
	}
	return result, nil
}

func (pr *productService) GetAllProducts(ctx context.Context) ([]entity.Product, error) {
	result, err := pr.productRepo.SelectAllProducts(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (pr *productService) UpdateProductById(ctx context.Context, productUpd entity.ProductUpdate, id uint64) (entity.Product, error) {
	product := entity.Product{}
	err := smapping.FillStruct(&product, smapping.MapFields(&productUpd))
	if err != nil {
		return entity.Product{}, err
	}
	result, err := pr.productRepo.UpdateProduct(ctx, product, id)
	if err != nil {
		return entity.Product{}, err
	}
	return result, nil
}

func (pr *productService) DeleteProductById(ctx context.Context, id uint64)  error{
	err := pr.productRepo.DeleteProduct(ctx, id);
	if err != nil {
		return err;
	}
	return nil
}
