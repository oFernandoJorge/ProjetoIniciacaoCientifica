package user

import(
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//UserHanlder é responsávl por receber requisições HTTP relacionados aos usuários
type Handler struct{
	service *Service
}

//NewUserHandler cria uma nova instância do UserHandler
func NewHandler(service *Service) *Handler{

	return &Handler{service: service}
}

//CreateUser recebe requisição via JSON para criar um novo usuário
func (h *Handler) Create(c *gin.Context){

	var user User

	//Faz bind do JSON recbido para struct User
	if err:= c.ShouldBindJSON(&user); err!= nil{
		
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Executa regras de negócio
	err := h.service.Create(&user)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

//GetAllUsers retorna todos os usuários cadastrados
func (h *Handler) FindAll(c *gin.Context){

	users, err := h.service.FindAll()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200,users)
}

//GetUserByID busca usuários pelo ID
func (h *Handler) FindByID(c *gin.Context){

	//Recupera ID da URL
	id,_ := strconv.Atoi(c.Param("id"))

	user, err := h.service.FindByID(uint(id))

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "Usuário não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}