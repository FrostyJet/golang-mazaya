package main

import (
	"encoding/json"
	"golang-mazaya/storefront/internal/db"
	"golang-mazaya/storefront/internal/dialer"
	"golang-mazaya/storefront/internal/domain"
	"golang-mazaya/storefront/internal/repository"
	"golang-mazaya/storefront/internal/service"
	"log"
)

func main() {
	log.Println("Storefront starting...")

	// Initialize database connection
	{
		err := db.Init()
		if err != nil {
			log.Panicf("Could not connect to database: %s\n", err.Error())
		}
		log.Println("Storefront connected to database successfully!")
	}

	recipeRepo := repository.NewRecipeRepository(db.DB)
	recipeService := service.NewRecipeService(recipeRepo)

	dialer.Subscribe("recipe.created", func(msg []byte) {
		log.Printf("Received a message: %s\n", msg)

		data := &domain.Recipe{}
		err := json.Unmarshal(msg, data)
		if err != nil {
			log.Printf("could parse message: %s", err.Error())
			return
		}

		recipe, err := recipeService.Create(data)
		if err != nil {
			log.Printf("could not create new recipe from %v\n", msg)
			log.Println(err.Error())
			return
		}

		log.Printf("New recipe has been created: %s\n", recipe.Title)
	})
}
