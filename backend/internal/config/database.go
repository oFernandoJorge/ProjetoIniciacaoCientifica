package config

import (
	"fmt"
	"log"
	"os"

	"ProjetoIniciacaoCientifica/internal/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB é a instância global do banco de dados
//Variável utilizada pelas camadas de repository
var DB *gorm.DB

//ConnectDatabase realiza:
//
//1. Carrega as variáveis de ambiente do arquivo .env
//2. Montagem da string de conexão (DSN) usando as variáveis de ambiente carregadas
//3. Estabelece a conexão com o banco de dados usando GORM
//4. Realiza a migração automática
func ConnectDatabase(){

	//Carrega as variáveis de ambiente do arquivo .env
	er := godotenv.Load()

	// Verifica se a variável de erro contém algum valor (se houve problema ao carregar o .env)
	if er != nil {
		log.Fatal("Error loading .env file")
	}

	//Recupera as variáveis de ambiente necessárias para a conexão com o banco de dados
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	//Montagem da string de conexão (DSN) usando as variáveis de ambiente carregadas
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host,
		user,
		password,
		dbname,
		port,
		sslmode,
	)

	//Estabelece a conexão com o banco de dados usando GORM
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	//Armazena conexão na variável global
	DB = database

	log.Println("Database connection established")

	Migrate()
}

//Migrate executa a criação automática das tabelas definidas nos models
func Migrate() {
	err := DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	log.Println("Database migrated successfully")
}