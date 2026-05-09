package routes

import(
	"ProjetoIniciacaoCientifica/internal/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configura as rotas da aplicação
func SetupRoutes(r *gin.Engine){

	userHandler := handlers.NewUserHandler()

	//Grupo de rotas para usuários
	users := r.Group("/users")
	{
		users.POST("/", userHandler.CreateUser)
		users.GET("/", userHandler.GetAllUsers)
		users.GET("/:id", userHandler.GetUserByID)
	}
}