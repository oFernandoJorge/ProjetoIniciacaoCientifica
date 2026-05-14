package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"ProjetoIniciacaoCientifica/internal/config"
	"ProjetoIniciacaoCientifica/internal/evaluator"
	"ProjetoIniciacaoCientifica/internal/pdf"
	"ProjetoIniciacaoCientifica/internal/room"
	"ProjetoIniciacaoCientifica/internal/session"
	"ProjetoIniciacaoCientifica/internal/submission"
	"ProjetoIniciacaoCientifica/internal/user"
)

func main() {

	// Inicializa conexão com banco
	config.ConnectDatabase()

	// =========================
	// MIGRATIONS
	// =========================

	// Migration do módulo user
	if err := user.Migrate(config.DB); err != nil {
		log.Println("[WARN] user migration failed:", err)
	}

	// Migration do módulo submission
	if err := submission.Migrate(config.DB); err != nil {
		log.Println("[WARN] submission migration failed:", err)
	}
	// Migration do módulo room
	if err := room.Migrate(config.DB); err != nil {
		log.Println("[WARN] room migration failed:", err)
	}
	// Migration do módulo evaluator
	if err := evaluator.Migrate(config.DB); err != nil {
		log.Println("[WARN] evaluator migration failed:", err)
	}
	// Migration do módulo session
	if err := session.Migrate(config.DB); err != nil {
		log.Println("[WARN] session migration failed:", err)
	}

	// =========================
	// GIN ROUTER
	// =========================

	r := gin.Default()
	r.Use(corsMiddleware())

	// Health check
	r.GET("/health", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "API funcionando",
		})
	})

	// =========================
	// USER MODULE
	// =========================

	userRepo := user.NewRepository()

	userService := user.NewService(userRepo)

	userHandler := user.NewHandler(userService)

	user.RegisterRoutes(r, userHandler)

	// =========================
	// SUBMISSION MODULE
	// =========================

	submissionRepo := submission.NewRepository()

	submissionService := submission.NewService(submissionRepo)

	submissionHandler := submission.NewHandler(submissionService)

	submission.RegisterRoutes(r, submissionHandler)

	// =========================
	// PDF MODULE
	// =========================

	pdfHandler := pdf.NewHandler()

	pdf.RegisterRoutes(r, pdfHandler)

	// =========================
	// ROOM MODULE
	// =========================

	roomRepo := room.NewRepository()

	roomService := room.NewService(roomRepo)

	// Gera salas automáticas
	if err := roomService.GenerateDefaultRooms(); err != nil {
		log.Println("[WARN] failed to generate default rooms:", err)
	}

	// =========================
	// EVALUATOR MODULE
	// =========================

	evaluatorRepo := evaluator.NewRepository()

	evaluatorService := evaluator.NewService(evaluatorRepo)

	evaluatorHandler := evaluator.NewHandler(evaluatorService)

	evaluator.RegisterRoutes(r, evaluatorHandler)

	// =========================
	// SESSION MODULE
	// =========================

	sessionRepo := session.NewRepository()

	sessionService := session.NewService(sessionRepo)

	_ = sessionService

	// =========================
	// START SERVER
	// =========================

	log.Println("Server running on port 8080")

	log.Fatal(r.Run(":8080"))
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
