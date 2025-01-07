package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
	"user/config"
	"user/internal/db/repositories"
	"user/internal/handlers"
	"user/internal/services"
)

func main() {
	logFile, err := os.OpenFile("logs/requests.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Ошибка открытия файла для логов: %v", err)
	}
	defer logFile.Close()

	logger := log.New(logFile, "", log.LstdFlags)
	cfg, err := config.FromEnv()
	if err != nil {
		log.Fatalf("Ошибка загрузки данных: %v", err)
		return
	}

	conn, err := sqlx.Connect("postgres", cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
		return
	}
	defer conn.Close()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{
		Logger:  logger,
		NoColor: true,
	}))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(5 * time.Second))

	userRepository := repositories.NewUserRepo(conn)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	r.Route("/user", func(r chi.Router) {
		r.Post("/login", userHandler.Login)
	})

	log.Println("Сервер запущен по адресу: ", cfg.ListenAddrAndPort())
	logger.Println("Сервер запущен по адресу: ", cfg.ListenAddrAndPort())

	err = http.ListenAndServe(cfg.ListenAddrAndPort(), r)
	if err != nil {
		logger.Fatalf("Ошибка запуска сервера: %v", err)
		return
	}

	log.Println("Сервер отключен")
	logger.Println("Сервер отключен")
}
