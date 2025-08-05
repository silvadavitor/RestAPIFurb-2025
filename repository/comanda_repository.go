package repository

import (
	"gorm.io/gorm"
)

// DAO|Repository

type ComandaRepository struct {
	db *gorm.DB
}

func NewComandaRepository(db *gorm.DB) *ComandaRepository {
	return &ComandaRepository{db: db}
}
