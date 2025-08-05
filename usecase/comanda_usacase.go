package usecase

import (
	"RestAPIFurb-2025/model"
	"RestAPIFurb-2025/repository"
	"errors"
)

// validacoes

type ComandaUsecase struct {
	repo *repository.ComandaRepository
}

func NewComandaUsecase(repo *repository.ComandaRepository) *ComandaUsecase {
	return &ComandaUsecase{repo: repo}
}

func (uc *ComandaUsecase) GetComandas() ([]model.Comanda, error) {
	return uc.repo.GetComandas()
}

func (uc *ComandaUsecase) GetComandaById(id uint) (model.Comanda, error) {
	return uc.repo.GetComandaById(id)
}

func (uc *ComandaUsecase) CreateComanda(comanda model.Comanda) (model.Comanda, error) {
	if len(comanda.Produtos) == 0 {
		return model.Comanda{}, errors.New("a comanda precisa ter ao menos um produto")
	}
	return uc.repo.CreateComanda(comanda)
}

func (uc *ComandaUsecase) UpdateComanda(id uint, comanda model.Comanda) (model.Comanda, error) {
	return uc.repo.UpdateComanda(id, comanda)
}

func (uc *ComandaUsecase) DeleteComanda(id uint) error {
	return uc.repo.DeleteComanda(id)
}
