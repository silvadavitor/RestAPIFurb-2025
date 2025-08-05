package usecase

import (
	"RestAPIFurb-2025/model"
	"RestAPIFurb-2025/repository"
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
	return uc.repo.CreateComanda(comanda)
}

func (uc *ComandaUsecase) UpdateComanda(id uint, comandaAtualizada model.Comanda) (model.Comanda, error) {
	// Buscar comanda atual no banco
	existente, err := uc.repo.GetComandaById(id)
	if err != nil {
		return model.Comanda{}, err
	}

	// Salvar comanda atualizada
	return uc.repo.UpdateComanda(id, existente)
}

func (uc *ComandaUsecase) DeleteComanda(id uint) error {
	return uc.repo.DeleteComanda(id)
}
