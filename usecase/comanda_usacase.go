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

func (uc *ComandaUsecase) UpdateComanda(id uint, comandaAtualizada model.Comanda) (model.Comanda, error) {
	// Buscar comanda atual no banco
	existente, err := uc.repo.GetComandaById(id)
	if err != nil {
		return model.Comanda{}, err
	}

	// Atualizar apenas os campos enviados (regra da prova)
	if comandaAtualizada.NomeUsuario != "" {
		existente.NomeUsuario = comandaAtualizada.NomeUsuario
	}
	if comandaAtualizada.TelefoneUsuario != "" {
		existente.TelefoneUsuario = comandaAtualizada.TelefoneUsuario
	}
	if comandaAtualizada.IDUsuario != 0 {
		existente.IDUsuario = comandaAtualizada.IDUsuario
	}

	// Só atualiza produtos se vier no JSON E tiver pelo menos 1 produto
	if comandaAtualizada.Produtos != nil && len(comandaAtualizada.Produtos) > 0 {
		existente.Produtos = comandaAtualizada.Produtos
	}

	// Salvar comanda atualizada
	return uc.repo.UpdateComanda(id, existente)
}

func (uc *ComandaUsecase) DeleteComanda(id uint) error {
	return uc.repo.DeleteComanda(id)
}
