package controller

import (
	"RestAPIFurb-2025/usecase"
)

type ComandaController struct {
	usecase *usecase.ComandaUsecase
}

func NewComandaController(uc *usecase.ComandaUsecase) *ComandaController {
	return &ComandaController{usecase: uc}
}
