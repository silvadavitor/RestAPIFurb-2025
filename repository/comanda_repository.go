package repository

import (
	"RestAPIFurb-2025/model"
	"errors"
	"fmt"
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
	err := repo.db.Find(&comandas).Error
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

func (repo *ComandaRepository) UpdateComanda(id uint, comanda model.Comanda) (model.Comanda, error) {
	err := repo.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&comanda).Error
	return comanda, err
}

func (repo *ComandaRepository) DeleteComanda(id uint) error {
	var comanda model.Comanda
	if err := repo.db.First(&comanda, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("comanda não encontrada")
		}
		return err
	}
	if err := repo.db.Where("comanda_id = ?", id).Delete(&model.Produto{}).Error; err != nil {
		return err
	}

	return repo.db.Delete(&comanda).Error
}
