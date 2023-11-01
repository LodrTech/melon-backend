package main

import (
	"log"
	"github.com/Marif226/melon/internal/app"
)

func main() {
	a, err := app.NewApp()
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}

	// router := chi.NewRouter()

	// s := service.NewProductService()
	// h := handler.NewProductHandler(
	// 	s,
	// )

	// router.Post("/product", h.Create)

	// http.ListenAndServe(":8000", router)
}