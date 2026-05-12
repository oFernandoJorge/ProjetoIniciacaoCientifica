package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Handler controla endpoints HTTP
type Handler struct {
	service *Service
}

// NewHandler cria novo handler
func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

// Create cria usuário
func (h *Handler) Create(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := h.service.Create(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "usuário criado com sucesso",
	})
}

// FindAll retorna usuários

func (h *Handler) FindAll(c *gin.Context) {
	users, err := h.service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, users)
}
