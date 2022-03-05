package main

import (
	"fmt"
	"os"
)

func main() {
	os.Exit(run())
}

func run() int {
	// ctx := context.Background()
	os.Setenv("TEST", "value")
	fmt.Printf("value: %s\n", os.Getenv("TEST"))

	return 0
}

// Handler
// func hello(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello, World!")
// }
