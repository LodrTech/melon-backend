package main

import (
	"context"
	"log"
	"github.com/Marif226/melon/internal/app"
)

func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatal("failed to init app: ", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatal("failed to run app: ", err)
	}
}