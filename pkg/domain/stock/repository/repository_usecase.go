package repository

import (
	"github.com/Mangaba-Labs/ape-finance-scrapper/pkg/domain/stock/model"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

// Create share in database
func (r *Repository) Create(share *model.Share) (err error) {
	result := r.DB.Create(&share)
	err = result.Error
	return
} 

// Delete share from database
func (r *Repository) Delete(ID int) (err error) {
	result := r.DB.Delete(&model.Share{}, ID)
	err = result.Error
	return
}

// FindAll shares in database
func (r *Repository) FindAll() (shares []model.Share, err error) {
	result := r.DB.Find(&shares)
	err = result.Error
	return
}

// FindByID shares in database
func (r *Repository) FindByID(ID int) (share model.Share, err error) {
	result := r.DB.First(&share, "id = ?", ID)
	err = result.Error
	return
}

// FindByBvmf shares in database
func (r *Repository) FindByBvmf(bvmf string) (share model.Share, err error) {
	result := r.DB.First(&share, "bvmf = ?", bvmf)
	err = result.Error
	return
}
