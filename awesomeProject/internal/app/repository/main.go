package repository

import (
	"awesomeProject/internal/app/ds"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(dsn string) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) GetProductByID(productId string) (*ds.Product, error) {
	product := &ds.Product{}

	err := r.db.First(product, "id = ?", productId).Error // find product with code D42
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (r *Repository) GetProductBelowID(productId string) ([]ds.Product, error) {
	var products []ds.Product

	var err error
	if productId == "" {
		err = r.db.Where("").Find(&products).Error // find product
	} else {
		err = r.db.Where("id < ? + 1", productId).Find(&products).Error
	}

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *Repository) CreateProduct(product ds.Product) error {
	return r.db.Create(product).Error
}
