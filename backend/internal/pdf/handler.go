package pdf

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

type PresentationPdfRow struct {
	Time          string `json:"time"`
	Title         string `json:"title"`
	PresenterName string `json:"presenterName"`
}

type PresentationPdfRoom struct {
	RoomName         string               `json:"roomName"`
	KnowledgeArea    string               `json:"knowledgeArea"`
	PresentationType string               `json:"presentationType"`
	Courses          []string             `json:"courses"`
	Submissions      []PresentationPdfRow `json:"submissions"`
}

type GeneratePresentationPdfRequest struct {
	Date   string                `json:"date"`
	Turn   string                `json:"turn"`
	Campus string                `json:"campus"`
	Items  []PresentationPdfRoom `json:"items"`
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Generate(c *gin.Context) {
	var payload GeneratePresentationPdfRequest

	contentType := c.GetHeader("Content-Type")
	if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") || strings.HasPrefix(contentType, "multipart/form-data") {
		raw := c.PostForm("payload")
		if raw == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "payload missing"})
			return
		}
		if err := json.Unmarshal([]byte(raw), &payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	buf := bytes.NewBuffer(nil)
	if err := GeneratePresentationSchedulePDF(payload, buf); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar PDF: " + err.Error()})
		return
	}

	if buf.Len() == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "O PDF gerado está vazio"})
		return
	}

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "inline; filename=organizador-de-salas.pdf")
	c.Header("Content-Length", fmt.Sprint(buf.Len()))

	c.Data(http.StatusOK, "application/pdf", buf.Bytes())
}
