package http

import (
	"app/internal/config"
	"app/internal/infrastructure/http/handlers"
	"app/internal/repository"
	"app/internal/usecases"
	"fmt"

	// jwtware "github.com/gofiber/contrib/jwt"
	_ "app/docs"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
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

// @title		EasyQuiz
// @version		0.0.1
// @description	Сервис для работы с Квизами
func (s *Server) Init(db *gorm.DB) {
	s.Server.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	userRepository := repository.NewUserRepository(db)
	userUseCase := usecases.NewUserUsecase(userRepository)
	userHandler := handlers.NewUserHandler(*userUseCase, s.Config)
	s.Server.Get("/metrics", monitor.New())
	cfg := swagger.Config{
		BasePath: "",
		FilePath: "./docs/swagger.json",
		Path:     "swagger",
		Title:    "Swagger API Docs",
	}
	s.Server.Use(swagger.New(cfg))
	s.Server.Post(
		"/login",
		userHandler.Login,
	)
	s.Server.Post(
		"/register",
		userHandler.Register,
	)
	// s.Server.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: jwtware.SigningKey{
	// 		JWTAlg: jwtware.RS256,
	// 		Key:    s.Config.HTTP.PublicKey,
	// 	},
	// 	SuccessHandler: handlers.JWTMiddleware,
	// }))
	quizRepository := repository.NewQuizRepository(db)
	resultRepository := repository.NewResultRepository(db)
	resultUseCase := usecases.NewResultUseCase(resultRepository)
	quizUseCase := usecases.NewQuizUseCase(quizRepository, *resultUseCase)
	quizHandler := handlers.NewQuizHandler(
		*quizUseCase,
		s.Config,
	)
	// создание квиза
	s.Server.Post(
		"/quiz",
		quizHandler.CreateQuiz,
	)
	// начать квиз
	s.Server.Post(
		"/quiz/:quiz_id/start",
		quizHandler.StartQuiz,
	)
	// создание вопроса для квиза
	s.Server.Post(
		"/quiz/:quiz_id/question",
		quizHandler.CreateQuestion,
	)
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
