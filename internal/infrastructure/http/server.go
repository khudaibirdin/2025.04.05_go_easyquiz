package http

import (
	"app/internal/infrastructure/http/handlers"
	"app/internal/repository"
	"app/internal/usecases"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Server struct {
	Server    fiber.App
	Host      string
	UrlPrefix string
}

func New(config config.Config) *Server {
	s := &Server{
		Server:    *fiber.New(),
		Host:      config.HTTP.Host,
		UrlPrefix: config.HTTP.Prefix,
	}
	return s
}

// Запуск сервера http
func (s *Server) Start() {
	err := s.Server.Listen(s.Host)
	if err != nil {
		panic(err)
	}
}

func (s *Server) Init(db *gorm.DB) {
	s.Server.Get(
		"/login",
		handlers.NewUserHandler(*usecases.NewUserUsecase(repository.NewUserRepository(db))).Login,
	)
}
