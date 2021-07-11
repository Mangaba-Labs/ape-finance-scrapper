package service

import (
	"github.com/Mangaba-Labs/ape-finance-scrapper/pkg/domain/models"
	"github.com/Mangaba-Labs/ape-finance-scrapper/pkg/domain/stock/model"
	stockRepository "github.com/Mangaba-Labs/ape-finance-scrapper/pkg/domain/stock/repository"
)

// StockService contract
type StockService interface {
	Create(bvmf string) models.Response
	Delete(ID int) models.Response
	GetAll() ([]model.Share, models.Response)
	GetByID(ID int) (model.Share, models.Response)
}

// NewUserService returns a StockService implementation
func NewUserService(repository stockRepository.Repository) (service StockService) {
	service = &Service{
		Repository: repository,
	}
	return
}
