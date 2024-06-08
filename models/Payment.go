package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Payment struct {
	gorm.Model
	ID          string `gorm:"type:char(36);primaryKey"`
	Gateway     string
	Amount      float64
	Currency    string
	Description string
	Status      string
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (payment *Payment) BeforeCreate(tx *gorm.DB) (err error) {
	payment.ID = uuid.New().String()
	return
}
