package user

import(
	"github.com/gin-gonic/gin"
)

// SetupRoutes configura as rotas da aplicação
func RegisterRoutes(r *gin.Engine, handler *Handler){

group := r.Group("/users")

group.POST("/", handler.Create)
group.GET("/", handler.FindAll)
group.GET("/:id", handler.FindByID)

}