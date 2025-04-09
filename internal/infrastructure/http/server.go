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
	Config *config.Config
}

func New(config *config.Config) *Server {
	s := &Server{
		Server: *fiber.New(),
		Config: config,
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
			Key:    s.Config.HTTP.PublicKey,
		},
	}))
	// создание квиза
	s.Server.Post(
		"/quiz",
		handlers.NewQuizHandler(*usecases.NewQuizUseCase(repository.NewQuizRepository(db)), s.Config).CreateQuiz,
	)
	// // создание вопроса для квиза
	// s.Server.Post(
	// 	"/quiz/:quiz_id/question",
	// )
	// // получение информации о квизах
	// s.Server.Get(
	// 	"/quiz",
	// )
	// // получение информации о конкретном квизе
	// s.Server.Get(
	// 	"/quiz/:quiz_id",
	// )
	// // получение вопросов по конкретному квизу
	// s.Server.Get(
	// 	"/quiz/:quiz_id/question",
	// )
	// // получение конкретного вопроса по конкретному квизу
	// s.Server.Get(
	// 	"/quiz/:quiz_id/question/:question_id",
	// )
}
