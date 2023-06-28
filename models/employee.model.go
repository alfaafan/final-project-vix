package models

import "gorm.io/gorm"

type Employees struct {
	gorm.Model
	SecretId uint
	Name string
	Email string
	Phone string
	Address string
	PositionsID uint 
	Positions Positions `gorm:"foreignKey:PositionsID"`
}