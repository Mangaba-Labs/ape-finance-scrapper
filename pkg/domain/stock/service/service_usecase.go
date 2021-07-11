package service

import (
	"github.com/Mangaba-Labs/ape-finance-scrapper/pkg/domain/models"
	"github.com/Mangaba-Labs/ape-finance-scrapper/pkg/domain/scrapper"
	"github.com/Mangaba-Labs/ape-finance-scrapper/pkg/domain/stock/model"
	"github.com/Mangaba-Labs/ape-finance-scrapper/pkg/domain/stock/repository"
)

// Service our implementation
type Service struct {
	Repository repository.Repository
}

// Create handle the creation of stock in database
func (s *Service) Create(bvmf string) (response models.Response) {
	share, err := s.Repository.FindByBvmf(bvmf)
	// if doesn't trigger error, the stock already exists in database
	if err == nil {
		response.Set(409, "Error", "Stock already exists!")
		return response
	}
	share, err = scrapper.ScrapFullStock(bvmf)
	if err != nil {
		response.Set(500, "Error", "Server internal error!")
		return response
	}
	err = s.Repository.Create(&share)

	if err != nil {
		response.Set(500, "Error", "Cannot create stock!")
		return response
	}
	response.Set(201, "Success", "Created!")
	return response
}

// Delete handle delete stock
func (s *Service) Delete(ID int) (response models.Response) {
	err := s.Repository.Delete(ID)

	if err != nil {
		response.Set(500, "Error", "Cannot delete stock!")
		return response
	}
	response.Set(204, "Success", "Deleted!")
	return response
}

// GetAll stocks in database
func (s *Service) GetAll() (stocks []model.Share, response models.Response) {
	stocks, err := s.Repository.FindAll()

	if err != nil {
		response.Set(500, "Error", "Cannot get stocks!")
		return nil, response
	}
	response.Set(200, "Success", "Ok")
	return stocks, response
}

// GetByID our stocks
func (s *Service) GetByID(ID int) (stock model.Share, response models.Response) {
	stock, err := s.Repository.FindByID(ID)

	if err != nil {
		response.Set(404, "Error", "Stock not found!")
		return stock, response
	}
	response.Set(200, "Success", "Ok")
	return stock, response
}
