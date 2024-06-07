package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	ID          uint `gorm:"primaryKey;autoIncrement"`
	Gateway     string
	Amount      float64
	Currency    string
	Description string
	Status      string
}
