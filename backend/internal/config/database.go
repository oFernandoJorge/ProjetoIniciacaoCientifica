package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB é a instância global do banco de dados
// utilizada pelos repositórios
var DB *gorm.DB

// ConnectDatabase realiza:
//
// 1. Carrega variáveis de ambiente (.env)
// 2. Monta string de conexão (DSN)
// 3. Conecta no PostgreSQL usando GORM
func ConnectDatabase() {

	// Carrega .env (opcional em produção)
	if err := godotenv.Load(); err != nil {
		log.Println("[WARN] .env not found, using system env vars")
	}

	// Variáveis de ambiente do banco
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	// DSN do PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host,
		user,
		password,
		dbname,
		port,
		sslmode,
	)

	// Conexão com banco via GORM
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("[ERROR] failed to connect database:", err)
	}

	DB = database

	log.Println("Database connection established")
}