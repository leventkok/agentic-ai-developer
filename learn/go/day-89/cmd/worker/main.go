package main

import (
	"context"
	"log"

	"learn/go/day-89/internal/worker"
)

func main() {
	if err := worker.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}
