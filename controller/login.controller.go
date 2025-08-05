package controller

import (
	"RestAPIFurb-2025/auth"
	"RestAPIFurb-2025/model"
	"RestAPIFurb-2025/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	usecase *usecase.LoginUsecase
}

func NewLoginController(uc *usecase.LoginUsecase) *LoginController {
	return &LoginController{usecase: uc}
}

// Login godoc
// @Summary Gera token JWT
// @Description Gera um token JWT simples para teste
// @Tags Autenticação
// @Accept json
// @Produce json
// @Param login body model.Login true "Dados do usuario"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /RestAPIFurb/login [post]
func (ctrl *LoginController) Login(ctx *gin.Context) {
	var login model.LoginInput

	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	if !ctrl.usecase.ValidarLogin(login) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha inválidos"})
		return
	}

	token, err := auth.GerarToken(login.Usuario)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
