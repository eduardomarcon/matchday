package http

import (
	"api/internal/adapter/storage/mongodb"
	"api/internal/adapter/storage/mongodb/repository"
	"api/internal/core/service"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int

	db mongodb.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	newServer := &Server{
		port: port,

		db: mongodb.New(),
	}

	dbHandler := NewHealthHandler(newServer.db)

	userRepo := repository.NewUserRepository(newServer.db.Get())
	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	router := NewRouter(*dbHandler, *userHandler)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.port),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
