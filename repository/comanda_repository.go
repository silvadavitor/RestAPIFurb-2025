package repository

import (
	"RestAPIFurb-2025/model"
	"gorm.io/gorm"
)

// DAO|Repository

type ComandaRepository struct {
	db *gorm.DB
}

func NewComandaRepository(db *gorm.DB) *ComandaRepository {
	return &ComandaRepository{db: db}
}

func (repo *ComandaRepository) GetComandas() ([]model.Comanda, error) {
	var comandas []model.Comanda
	err := repo.db.Preload("Produtos").Find(&comandas).Error
	return comandas, err
}

func (repo *ComandaRepository) GetComandaById(id uint) (model.Comanda, error) {
	var comanda model.Comanda
	err := repo.db.Preload("Produtos").First(&comanda, id).Error
	return comanda, err
}

func (repo *ComandaRepository) CreateComanda(comanda model.Comanda) (model.Comanda, error) {
	err := repo.db.Create(&comanda).Error
	return comanda, err
}
