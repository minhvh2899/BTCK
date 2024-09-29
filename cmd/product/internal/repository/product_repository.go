package repository

import (
	"my-project/cmd/product/internal/models"

	"gorm.io/gorm"
)



type ProductRepository interface {
    Create(product *models.Product) (*models.Product, error)
    FindByID(id uint) (*models.Product, error)
    Update(product *models.Product) (*models.Product, error)
    Delete(id uint) error
    FindAll() ([]models.Product, error)
}

type productRepository struct {
    db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
    return &productRepository{db: db}
}

func (r *productRepository) Create(product *models.Product) (*models.Product, error) {
    result := r.db.Create(product)
    return product, result.Error
}

func (r *productRepository) FindByID(id uint) (*models.Product, error) {
    var product models.Product
    result := r.db.First(&product, id)
    return &product, result.Error
}

func (r *productRepository) Update(product *models.Product) (*models.Product, error) {
    result := r.db.Save(product)
    return product, result.Error
}

func (r *productRepository) Delete(id uint) error {
    result := r.db.Delete(&models.Product{}, id)
    return result.Error
}

func (r *productRepository) FindAll() ([]models.Product, error) {
    var products []models.Product
    result := r.db.Find(&products)
    return products, result.Error
}