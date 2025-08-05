package model

type Produto struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	Nome      string  `json:"nome" binding:"required"`
	Preco     float64 `json:"preco" binding:"required,gt=0"`
	ComandaID uint    `json:"-"`
}
