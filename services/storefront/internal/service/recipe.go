package service

import (
	"golang-mazaya/storefront/internal/domain"
	"log"
	"sync"
)

var once sync.Once

type recipeService struct {
	recipeRepository domain.RecipeRepository
}

var instance *recipeService

func NewRecipeService(r domain.RecipeRepository) domain.RecipeService {
	once.Do(func() {
		instance = &recipeService{
			recipeRepository: r,
		}
	})

	return instance
}

func (*recipeService) Save() {
	log.Println("Insert has been called!")
}
