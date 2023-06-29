package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	transactionType string
	amount uint64
	note string
}