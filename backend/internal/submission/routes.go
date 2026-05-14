package submission

import "github.com/gin-gonic/gin"

// RegisterRoutes registra rotas do módulo
func RegisterRoutes(r *gin.Engine, handler *Handler) {

	submissions := r.Group("/submissions")

	{
		submissions.POST("/", handler.Create)
		submissions.GET("/", handler.FindAll)
	}
}