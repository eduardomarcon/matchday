package main

import (
	"api/internal/adapter/handler/http"
	"api/internal/adapter/logger"
	"fmt"
)

func main() {
	logger.Set()

	server := http.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
