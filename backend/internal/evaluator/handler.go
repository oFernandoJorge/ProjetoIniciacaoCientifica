package evaluator

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler controla endpoints
type Handler struct {
	service *Service
}

// NewHandler cria handler
func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

// Create cria avaliador
func (h *Handler) Create(c *gin.Context) {

	var dto CreateEvaluatorDTO

	if err := c.ShouldBindJSON(&dto); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	evaluator := Evaluator{
		Name: dto.Name,

		Email: dto.Email,

		Course: dto.Course,

		KnowledgeArea: dto.KnowledgeArea,

		AvailableMorning: dto.AvailableMorning,

		AvailableAfternoon: dto.AvailableAfternoon,

		AvailableNight: dto.AvailableNight,

		MaxPresentations: dto.MaxPresentations,

		AcceptedPresentationType: dto.AcceptedPresentationType,
	}

	err := h.service.Create(&evaluator)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, evaluator)
}

// FindAll retorna avaliadores
func (h *Handler) FindAll(c *gin.Context) {

	evaluators, err := h.service.FindAll()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, evaluators)
}