package repository

import "RestAPIFurb-2025/model"

// Interface
type LoginRepository interface {
	BuscarUsuario(login model.LoginInput) bool
}

// Implementação mockada
type MockLoginRepository struct{}

func NewLoginRepository(_ interface{}) LoginRepository {
	// O parâmetro (como dbConnection) é ignorado
	return &MockLoginRepository{}
}

func (r *MockLoginRepository) BuscarUsuario(login model.LoginInput) bool {
	// Mock
	return login.Usuario == "admin" && login.Senha == "123456"
}
