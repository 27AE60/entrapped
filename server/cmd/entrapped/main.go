package main

import (
	"github.com/SKatiyar/entrapped/server"
	"os"
)

func main() {
	envPort := os.Getenv("PORT")
	if len(envPort) != 0 {
		envPort = ":" + envPort
	}

	entrapped.Start(envPort)
}
