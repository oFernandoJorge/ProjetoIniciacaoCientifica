package user

import "github.com/gin-gonic/gin"

// RegisterRoutes registra rotas do módulo user
func RegisterRoutes(r *gin.Engine, handler *Handler) {
	users := r.Group("/users")
	{
		users.POST("/", handler.Create)
		users.GET("/", handler.FindAll)
	}
}
