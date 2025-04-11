package config

import "gorm.io/gorm"

var (
	db *gorm.DB
	logger *Logger
)


func Init() error {
	var err error
	
	db, err = InitializePostgres()
	if err != nil {
		logger.Errorf("Error initializing postgres: %v", err)
		return err
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = newLogger(p)
	return logger
}

