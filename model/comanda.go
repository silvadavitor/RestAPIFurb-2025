package model

type Comanda struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	IDUsuario       uint      `json:"idUsuario" binding:"required"`
	NomeUsuario     string    `json:"nomeUsuario" binding:"required"`
	TelefoneUsuario string    `json:"telefoneUsuario" binding:"required"`
	Produtos        []Produto `gorm:"foreignKey:ComandaID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"produtos" binding:"required,dive"`
}
