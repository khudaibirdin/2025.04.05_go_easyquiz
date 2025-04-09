package http

import (
	"app/internal/config"
	"app/internal/infrastructure/http/handlers"
	"app/internal/repository"
	"app/internal/usecases"
	"fmt"

	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/contrib/jwt"
	"gorm.io/gorm"
)

type Server struct {
	Server fiber.App
	Config    *config.Config
}

func New(config *config.Config) *Server {
	s := &Server{
		Server: *fiber.New(),
		Config:    config,
	}
	return s
}

// Запуск сервера http
func (s *Server) Start() {
	err := s.Server.Listen(fmt.Sprintf("%s:%s", s.Config.HTTP.Host, s.Config.HTTP.Port))
	if err != nil {
		panic(err)
	}
}

func (s *Server) Init(db *gorm.DB) {
	s.Server.Post(
		"/login",
		handlers.NewUserHandler(*usecases.NewUserUsecase(repository.NewUserRepository(db)), s.Config).Login,
	)
	s.Server.Post(
		"/register",
		handlers.NewUserHandler(*usecases.NewUserUsecase(repository.NewUserRepository(db)), s.Config).Register,
	)
	s.Server.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.RS256,
			Key:    []byte(s.Config.HTTP.JWTKey),
		},
	}))
	// s.Server.Post(
	// 	"/quiz",
	// 	handlers.NewUserHandler(*usecases.NewQuizUseCase(repository.NewQuizRepository(db)), s.Config).CreateQuiz,
	// )
}
