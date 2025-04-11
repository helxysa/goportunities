package config

import (
	"github.com/helxysa/goportunities.git/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	logger := GetLogger(">> POSTGRES: ")
	logger.Info("Initializing postgres...")
	//Criando database e conectando
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=123@senha dbname=go-testando port=5432 sslmode=disable TimeZone=America/Sao_Paulo"), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error connecting to database: %v", err)
		return nil, err
	}
	//Vendo se tem algum erro na migração
	err = db.AutoMigrate( &schemas.Opening{})
	if err != nil {
		logger.Errorf("Error migrating database: %v", err)
		return nil, err
	}

	//Se não deu nenhum erro então faz a migração
	return db, nil
	
}