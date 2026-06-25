package server

import (
	"go-starter-backend/internal/config"
	"go-starter-backend/internal/middleware"
	"go-starter-backend/internal/user"
	"go-starter-backend/pkg/response"
	"go-starter-backend/pkg/storage"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	cfg     *config.Config
	db      *sqlx.DB
	storage *storage.Storage
}

func New(cfg *config.Config, db *sqlx.DB, storage *storage.Storage) *Server {
	return &Server{
		cfg:     cfg,
		db:      db,
		storage: storage,
	}
}

func (s *Server) Routes() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CORS())
	router.GET("/health", s.health)

	userRepo := user.NewRepository(s.db)
	userService := user.NewService(userRepo, s.cfg)
	userHandler := user.NewHandler(userService)

	api := router.Group("/api/v1")
	user.RegisterRoutes(api, userHandler, middleware.Auth(s.cfg))

	return router
}

func (s *Server) health(c *gin.Context) {

	response.OK(c, "Server Healthy", gin.H{
		"status": "ok",
	})
}
