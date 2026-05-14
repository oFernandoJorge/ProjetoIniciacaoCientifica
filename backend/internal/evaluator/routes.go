package evaluator

import "github.com/gin-gonic/gin"

// RegisterRoutes registra rotas
func RegisterRoutes(r *gin.Engine, handler *Handler) {

	evaluators := r.Group("/evaluators")

	{
		evaluators.POST("/", handler.Create)
		evaluators.GET("/", handler.FindAll)
	}
}