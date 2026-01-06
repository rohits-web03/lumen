package main

import (
	"log"

	"github.com/rohits-web03/lumen/internal/server"
)

func main() {
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
