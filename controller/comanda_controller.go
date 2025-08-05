package controller

import (
	"RestAPIFurb-2025/model"
	"RestAPIFurb-2025/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ComandaController struct {
	usecase *usecase.ComandaUsecase
}

func NewComandaController(uc *usecase.ComandaUsecase) *ComandaController {
	return &ComandaController{usecase: uc}
}

func (ctrl *ComandaController) GetComandas(ctx *gin.Context) {
	comandas, err := ctrl.usecase.GetComandas()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, comandas)
}

func (ctrl *ComandaController) GetComandaById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	comanda, err := ctrl.usecase.GetComandaById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Comanda não encontrada"})
		return
	}
	ctx.JSON(http.StatusOK, comanda)
}

func (ctrl *ComandaController) CreateComanda(ctx *gin.Context) {
	var comanda model.Comanda

	if err := ctx.ShouldBindJSON(&comanda); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comandaCriada, err := ctrl.usecase.CreateComanda(comanda)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, comandaCriada)
}
