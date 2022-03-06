package main

import (
	"log"
	"os"

	"github.com/k3forx/coffee_memo/pkg/config"
	"github.com/k3forx/coffee_memo/pkg/inject"
	"github.com/k3forx/coffee_memo/pkg/logger"
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

	fn, err := logger.Init()
	if err != nil {
		log.Fatalf("init logger failed: %+v\n", err)
	}
	defer fn()

	injector, fn, err := inject.New()
	if err != nil {
		log.Fatalln(err)
		return 1
	}
	defer fn()

	s := server.NewServer(injector)
	if err := s.Start(); err != nil {
		log.Println(err)
		return 1
	}

	return 0
}
