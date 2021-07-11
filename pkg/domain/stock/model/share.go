package model

import "gorm.io/gorm"

type Share struct {
	gorm.Model
	Bvmf      string  `gorm:"not null" json:"bvmf"`
	Company   string  `gorm:"not null" json:"company"`
	Image     string  `gorm:"not null" json:"image"`
	Price     float32 `gorm:"not null" json:"price"`
	Variation string  `gorm:"not null" json:"variation"`
}

type VariableData struct {
	Variation string
	Price     float32
}