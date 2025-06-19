package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	DB_HOST          string
	DB_NAME          string
	DB_USER          string
	DB_PASSWORD      string
	DB_ROOT_PASSWORD string
	DB_PORT          string
	DB_SCHEMA        string
	connection       *gorm.DB
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		DB_HOST:          "localhost",
		DB_NAME:          "house",
		DB_USER:          "user",
		DB_PASSWORD:      "123",
		DB_ROOT_PASSWORD: "root",
		DB_PORT:          "3327",
		DB_SCHEMA:        "historico_consumo",
		connection:       nil,
	}
}

func (db_config *DBConfig) connectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", db_config.DB_USER, db_config.DB_PASSWORD, db_config.DB_HOST, db_config.DB_PORT, db_config.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func (db_config *DBConfig) GetConnection() *gorm.DB {
	if db_config.connection == nil {
		conn, err := db_config.connectDB()
		if err != nil {
			panic("failed to connect database")
		}
		db_config.connection = conn
	}
	return db_config.connection
}
