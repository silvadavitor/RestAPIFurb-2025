package model

type LoginInput struct {
	Usuario string `json:"usuario" example:"admin"`
	Senha   string `json:"senha" example:"123456"`
}
