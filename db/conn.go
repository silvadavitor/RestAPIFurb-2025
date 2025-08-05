package db

import (
	"RestAPIFurb-2025/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "go-db"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "postgres"
)

func ConnectDB() (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.Comanda{}, &model.Produto{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Conectado ao banco com GORM e tabelas migradas!")
	return db, nil
}
