package usecase

import (
	"RestAPIFurb-2025/repository"
)

// validacoes

type ComandaUsecase struct {
	repo *repository.ComandaRepository
}

func NewComandaUsecase(repo *repository.ComandaRepository) *ComandaUsecase {
	return &ComandaUsecase{repo: repo}
}
