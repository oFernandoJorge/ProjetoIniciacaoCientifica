package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
	"ProjetoIniciacaoCientifica/internal/config"
	"ProjetoIniciacaoCientifica/internal/routes"
)

// @title Projeto Iniciação Científica API
// @version 1.0
// @description API para gerenciamento de usuários e projetos de iniciação científica.
// @host localhost:8080
// @BasePath /
func main() {

	//Inicializa conexão com o banco de dados
	config.ConnectDatabase()

	//Cria instancia do router Gin
	r := gin.Default()

	//Endpoint de teste para verificar se a API está funcionando
	// @Summary Verificar saúde da API
	// @Description Retorna uma mensagem de sucesso se a API estiver funcionando
	// @Tags Health
	// @Accept json
	// @Produce json
	// @Success 200 {object} map[string]string
	// @Router /health [get]
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API funcionando",
		})
	})

	//Registra rotas da aplicação
	routes.SetupRoutes(r)

	//Inicia o servidor na porta 8080
	r.Run(":8080")
}
