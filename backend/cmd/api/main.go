package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
	"ProjetoIniciacaoCientifica/internal/config"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API funcionando",
		})
	})

	r.Run(":8080")
}
