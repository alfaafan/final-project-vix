package models

import (
	"gorm.io/gorm"
)

type Positions struct {
	gorm.Model
	Name      string
	Salary    uint64
}


