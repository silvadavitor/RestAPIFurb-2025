package usecase

import (
	"RestAPIFurb-2025/model"
	"RestAPIFurb-2025/repository"
)

type LoginUsecase struct {
	repo repository.LoginRepository
}

func NewLoginUsecase(repo repository.LoginRepository) *LoginUsecase {
	return &LoginUsecase{repo: repo}
}

func (uc *LoginUsecase) ValidarLogin(login model.LoginInput) bool {
	return uc.repo.BuscarUsuario(login)
}
