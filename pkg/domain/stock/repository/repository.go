package repository

import (
	"github.com/Mangaba-Labs/ape-finance-scrapper/pkg/domain/stock/model"
	"gorm.io/gorm"
)

// StockRepository Contract
type StockRepository interface {
	Create(*model.Share) (error)
	FindAll() ([]model.Share,error)
	FindByID(ID int) (model.Share, error)
	FindByBvmf(bvmf string) (model.Share, error)
	Delete(ID int) (error)
}

// NewStockRepository repository postgres implementation
func NewStockRepository(db *gorm.DB) StockRepository {
	return &Repository{
		DB: db,
	}
}