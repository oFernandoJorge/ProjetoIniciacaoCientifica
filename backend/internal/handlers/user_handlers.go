package handlers

import(
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ProjetoIniciacaoCientifica/internal/models"
	"ProjetoIniciacaoCientifica/internal/services"
)

//UserHanlder é responsávl por receber requisições HTTP relacionados aos usuários
type UserHandler struct{
	service *services.UserService
}

//NewUserHandler cria uma nova instância do UserHandler
func NewUserHandler() *UserHandler{

	return &UserHandler{
		service: services.NewUserService(),
	}
}

//CreateUser recebe requisição via JSON para criar um novo usuário
func (h *UserHandler) CreateUser(c *gin.Context){

	var user models.User

	//Faz bind do JSON recbido para struct User
	if err:= c.ShouldBindJSON(&user); err!= nil{
		
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados de usuário inválidos",
		})

		return
	}

	//Executa regras de negócio
	err := h.service.CreateUser(&user)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuário criado com sucesso",
		"user": user,
	})
}

//GetAllUsers retorna todos os usuários cadastrados
func (h *UserHandler) GetAllUsers(c *gin.Context){

	users, err := h.service.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

//GetUserByID busca usuários pelo ID
func (h *UserHandler) GetUserByID(c *gin.Context){

	//Recupera ID da URL
	idParam := c.Param("id")

	//Converte string para uint
	id, err := strconv.ParseUint(idParam,10,32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de usuário inválido",
		})
		return
	}

	user, err := h.service.GetUserByID(uint(id))

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Usuário não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}