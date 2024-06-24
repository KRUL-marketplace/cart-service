package main

import (
	"context"
	"github.com/KRUL-marketplace/cart-service/internal/app"
	"log"
)

func main() {
	ctx := context.Background()
	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run: %s", err.Error())
	}
}
