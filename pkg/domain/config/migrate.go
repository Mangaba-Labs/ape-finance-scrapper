package config

import (
	"log"

	"github.com/Mangaba-Labs/ape-finance-scrapper/pkg/domain/stock/model"
	"gorm.io/gorm"
)

// Migrate struct for database
type Migrate struct {
	DB *gorm.DB
}

// MigrateAll database migration
func (m *Migrate) MigrateAll() (err error) {
	log.Println("Migrating database... 🤞")
	err = m.DB.AutoMigrate(&model.Share{})

	if err != nil {
		log.Fatal("Something went wrong on db migration process...\n ", err)
	}

	log.Println("Database migrated with success 😁")
	return
}
