package repository

import (
	"context"
	"hacktiv8_fp_2/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	InsertNewProduct(ctx context.Context, product entity.Product) (entity.Product, error)
	SelectAllProducts(ctx context.Context) ([]entity.Product, error)
	UpdateProduct(ctx context.Context, product entity.Product, id uint64) (entity.Product, error)
	DeleteProduct(ctx context.Context, id uint64) error
}

type productRepository struct {
	connection *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepository {
	return &productRepository{connection: db}
}

func (pr *productRepository) InsertNewProduct(ctx context.Context, product entity.Product) (entity.Product, error){
	tx := pr.connection.Create(&product)
	if tx.Error != nil {
		return entity.Product{}, tx.Error
	}
	return product, nil
}
func (pr *productRepository) SelectAllProducts(ctx context.Context) ([]entity.Product, error){
	products := []entity.Product{}
	tx := pr.connection.Preload("User").Find(&products)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return products, nil
}
func (pr *productRepository) UpdateProduct(ctx context.Context, product entity.Product, id uint64) (entity.Product, error){
	tx := pr.connection.Save(&product)
	if tx.Error != nil {
		return entity.Product{}, tx.Error
	}
	return product, nil
}
func (pr *productRepository) DeleteProduct(ctx context.Context, id uint64) error {
	tx := pr.connection.Delete(entity.Product{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

