package submission

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler controla endpoints do módulo
type Handler struct {
	service *Service
}

// NewHandler cria handler
func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

// Create cria submissão
func (h *Handler) Create(c *gin.Context) {

	var submission Submission

	if err := c.ShouldBindJSON(&submission); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	err := h.service.Create(&submission)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "submissão criada com sucesso",
	})
}

// FindAll retorna submissões
func (h *Handler) FindAll(c *gin.Context) {

	submissions, err := h.service.FindAll()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, submissions)
}