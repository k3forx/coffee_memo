package main

import (
	"log"
	"os"

	"github.com/k3forx/coffee_memo/pkg/config"
	"github.com/k3forx/coffee_memo/pkg/server"
)

func main() {
	os.Exit(run())
}

func run() int {
	// ctx := context.Background()

	if err := config.LoadConfig(); err != nil {
		log.Println(err)
		return 1
	}

	s := server.NewServer()
	if err := s.Start(); err != nil {
		log.Println(err)
		return 1
	}

	return 0
}

// Handler
// func hello(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello, World!")
// }
