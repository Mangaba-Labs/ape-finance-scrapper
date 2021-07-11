package service

import (
	"github.com/Mangaba-Labs/ape-finance-scrapper/pkg/domain/models"
	"github.com/Mangaba-Labs/ape-finance-scrapper/pkg/domain/stock/model"
	"github.com/Mangaba-Labs/ape-finance-scrapper/pkg/domain/stock/repository"
)

type Service struct {
	Repository repository.Repository
}

func (s *Service) Create(bvmf string) (response models.Response) {
	var share model.Share
	err := s.Repository.Create(&share)

	if err != nil {
		response.Set(500, "Error", "Cannot create stock!")
		return response
	}
	response.Set(201, "Success","Created!")
	return response
}

func (s *Service) Delete(ID int) (response models.Response) {
	err := s.Repository.Delete(ID)

	if err != nil {
		response.Set(500, "Error", "Cannot delete stock!")
		return response
	}
	response.Set(204, "Success", "Deleted!")
	return response
}

func (s *Service) GetAll() (stocks []model.Share, response models.Response) {
	stocks, err := s.Repository.FindAll()

	if err != nil {
		response.Set(500, "Error","Cannot get stocks!")
		return nil, response
	}
	response.Set(200, "Success", "Ok")
	return stocks, response
}

func (s *Service) GetByID(ID int) (stock model.Share, response models.Response) {
	stock, err := s.Repository.FindByID(ID)

	if err != nil {
		response.Set(500, "Error", "Cannot get stock!")
		return stock, response
	}
	response.Set(200, "Success", "Ok")
	return stock, response
}
