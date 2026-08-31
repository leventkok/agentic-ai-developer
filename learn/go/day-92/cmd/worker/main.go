package main

import (
	"context"
	"log"

	"learn/go/day-92/internal/worker"
)

func main() {
	if err := worker.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}
