package server

import (
	"go-starter-backend/internal/config"
	"go-starter-backend/internal/user"
	"go-starter-backend/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	cfg *config.Config
	db  *sqlx.DB
}

func New(cfg *config.Config, db *sqlx.DB) *Server {
	return &Server{
		cfg: cfg,
		db:  db,
	}
}

func (s *Server) Routes() *gin.Engine {
	router := gin.Default()

	router.GET("/health", s.health)

	userRepo := user.NewRepository(s.db)
	userService := user.NewService(userRepo, s.cfg)
	userHandler := user.NewHandler(userService)

	api := router.Group("/api/v1")
	user.RegisterRoutes(api, userHandler)

	return router
}

func (s *Server) health(c *gin.Context) {

	response.OK(c, "Server Healthy", gin.H{
		"status": "ok",
	})
}
