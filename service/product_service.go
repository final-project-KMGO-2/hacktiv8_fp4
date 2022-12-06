package service

import (
	"context"
	"hacktiv8_fp_2/entity"
	"hacktiv8_fp_2/repository"
)

type ProductService interface {
	CreateNewProduct(ctx context.Context, productCrt entity.ProductCreate)
	GetAllProducts(ctx context.Context)
	UpdateProductById(ctx context.Context, productUpd entity.ProductUpdate, id uint64)
	DeleteProductById(ctx context.Context, id uint64)
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(pr repository.ProductRepository) ProductService {
	return &productService{productRepo: pr}
}

func (pr *productService) CreateNewProduct(ctx context.Context, productCrt entity.ProductCreate) {}
func (pr *productService) GetAllProducts(ctx context.Context) {}
func (pr *productService) UpdateProductById(ctx context.Context, productUpd entity.ProductUpdate, id uint64){}
func (pr *productService) DeleteProductById(ctx context.Context, id uint64){}

