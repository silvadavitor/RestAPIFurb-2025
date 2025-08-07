package controller

import (
	"RestAPIFurb-2025/model"
	"RestAPIFurb-2025/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ComandaController struct {
	usecase *usecase.ComandaUsecase
}

func NewComandaController(uc *usecase.ComandaUsecase) *ComandaController {
	return &ComandaController{usecase: uc}
}

// GetComandas godoc
// @Summary Lista todas as comandas
// @Description Retorna todas as comandas cadastradas
// @Tags Comandas
// @Produce json
// @Success 200 {array} model.Comanda
// @Router /RestAPIFurb/comandas [get]
func (ctrl *ComandaController) GetComandas(ctx *gin.Context) {
	comandas, err := ctrl.usecase.GetComandas()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var dto []model.ComandaResumoDTO
	for _, c := range comandas {
		dto = append(dto, model.ComandaResumoDTO{
			IDUsuario:       c.IDUsuario,
			NomeUsuario:     c.NomeUsuario,
			TelefoneUsuario: c.TelefoneUsuario,
		})
	}
	ctx.JSON(http.StatusOK, dto)
}

// GetComandaById godoc
// @Summary Busca comanda por ID
// @Description Retorna uma comanda pelo ID
// @Tags Comandas
// @Produce json
// @Param id path int true "ID da Comanda"
// @Success 200 {object} model.Comanda
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /RestAPIFurb/comandas/{id} [get]
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

// CreateComanda godoc
// @Summary Cria nova comanda
// @Description Cria uma nova comanda com os dados fornecidos
// @Tags Comandas
// @Accept json
// @Produce json
// @Param comanda body model.Comanda true "Dados da nova comanda"
// @Success 201 {object} model.Comanda
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /RestAPIFurb/comandas [post]
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

// UpdateComanda godoc
// @Summary Atualiza comanda existente
// @Description Atualiza uma comanda existente pelo ID
// @Tags Comandas
// @Accept json
// @Produce json
// @Param id path int true "ID da Comanda"
// @Param comanda body model.Comanda true "Dados para atualizar a comanda"
// @Success 200 {object} model.Comanda
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /RestAPIFurb/comandas/{id} [put]
func (ctrl *ComandaController) UpdateComanda(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var payload model.ComandaUpdateDTO
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Busca a comanda atual
	comandaAtual, err := ctrl.usecase.GetComandaById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Comanda não encontrada"})
		return
	}

	// Atualiza apenas os campos presentes
	if payload.IDUsuario != nil {
		comandaAtual.IDUsuario = *payload.IDUsuario
	}
	if payload.NomeUsuario != nil {
		comandaAtual.NomeUsuario = *payload.NomeUsuario
	}
	if payload.TelefoneUsuario != nil {
		comandaAtual.TelefoneUsuario = *payload.TelefoneUsuario
	}
	if payload.Produtos != nil {
		var produtos []model.Produto
		for _, p := range *payload.Produtos {
			produtos = append(produtos, model.Produto{
				ID:    p.ID,
				Nome:  p.Nome,
				Preco: p.Preco,
			})
		}
		comandaAtual.Produtos = produtos
	}

	comandaAtualizada, err := ctrl.usecase.UpdateComanda(uint(id), comandaAtual)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, comandaAtualizada)
}

// DeleteComanda godoc
// @Summary Remove comanda existente
// @Description Remove uma comanda existente pelo ID (Autenticação JWT necessária)
// @Tags Comandas
// @Security ApiKeyAuth
// @Param id path int true "ID da Comanda"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /RestAPIFurb/comandas/{id} [delete]
func (ctrl *ComandaController) DeleteComanda(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = ctrl.usecase.DeleteComanda(uint(id))
	if err != nil {
		if err.Error() == "comanda nao encontrada" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar comanda"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": gin.H{"text": "comanda removida"}})
}

// DummyModelDoc godoc
// @Tags Docs
// @Summary Modelos usados na API
// @Description Modelos auxiliares para documentação do Swagger
// @Success 200 {object} model.Comanda
// @Success 200 {object} model.ComandaResumoDTO
// @Success 200 {object} model.ComandaUpdateDTO
// @Success 200 {object} model.Produto
// @Success 200 {object} model.ProdutoDTO
// @Success 200 {object} model.LoginInput
// @Router /RestAPIFurb/docs/models [get]
func DummyModelDoc(ctx *gin.Context) {}
