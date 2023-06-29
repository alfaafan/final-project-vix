package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Type string
	Amount uint64
	Note string
	CompaniesID uint 
	Company Company `gorm:"foreignKey:CompaniesID"`
}