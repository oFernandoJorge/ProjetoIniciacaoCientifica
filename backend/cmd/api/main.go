package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"ProjetoIniciacaoCientifica/internal/config"
	"ProjetoIniciacaoCientifica/internal/user"
)

// @title Projeto Iniciação Científica API
// @version 1.0
// @description API para gerenciamento de usuários e projetos de iniciação científica.
// @host localhost:8080
// @BasePath /
func main() {

	// Inicializa conexão com o banco de dados
	config.ConnectDatabase()

	// 🔥 migrations por módulo (User é dono do próprio schema)
	if err := user.Migrate(config.DB); err != nil {
		log.Println("[WARN] user migration failed:", err)
	}

	// Cria instância do router Gin
	r := gin.Default()

	// Health check da API
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API funcionando",
		})
	})

	// Dependency Injection (User Module)
	repo := user.NewRepository()
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	// Registra rotas do módulo User
	user.RegisterRoutes(r, handler)

	// Inicia servidor HTTP
	log.Fatal(r.Run(":8080"))
}