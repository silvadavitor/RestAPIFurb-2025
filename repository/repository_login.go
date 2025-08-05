package repository

import "RestAPIFurb-2025/model"

// Interface
type LoginRepository interface {
	BuscarUsuario(login model.LoginInput) bool
}

// Implementação mockada
type MockLoginRepository struct{}

func NewLoginRepository(_ interface{}) LoginRepository {
	// O parâmetro (como dbConnection) é ignorado nesse mock
	return &MockLoginRepository{}
}

func (r *MockLoginRepository) BuscarUsuario(login model.LoginInput) bool {
	// Mock: simula validação de usuário/senha no “banco”
	return login.Usuario == "admin" && login.Senha == "123456"
}
