package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
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
	// Carrega variáveis do arquivo .env
	err := godotenv.Load("../.env")
	if err != nil {
		panic("err")
		fmt.Println("Erro ao carregar o arquivo .env")
	}
	return &DBConfig{
		DB_HOST:          os.Getenv("DB_HOST"),
		DB_NAME:          os.Getenv("DB_NAME"),
		DB_USER:          os.Getenv("DB_USER"),
		DB_PASSWORD:      os.Getenv("DB_PASSWORD"),
		DB_ROOT_PASSWORD: os.Getenv("DB_ROOT_PASSWORD"),
		DB_PORT:          os.Getenv("DB_PORT"),
		DB_SCHEMA:        os.Getenv("DB_SCHEMA"),
		connection:       nil,
	}
}

func (db_config *DBConfig) connectDB() (*gorm.DB, error) {
	fmt.Println(db_config.DB_HOST)
	fmt.Println("------------------------")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		db_config.DB_USER,
		db_config.DB_PASSWORD,
		db_config.DB_HOST,
		db_config.DB_PORT,
		db_config.DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func (db_config *DBConfig) GetConnection() *gorm.DB {
	if db_config.connection == nil {
		conn, err := db_config.connectDB()
		if err != nil {
			panic("failed to connect database: " + err.Error())
		}
		db_config.connection = conn
	}
	return db_config.connection
}
