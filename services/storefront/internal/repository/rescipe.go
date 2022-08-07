package repository

import (
	"golang-mazaya/storefront/internal/domain"
	"sync"
)

var once sync.Once

type recipeRepository struct{}

var instance domain.RecipeRepository

func NewRecipeRepository() domain.RecipeRepository {
	once.Do(func() {
		instance = &recipeRepository{}
	})

	return instance
}
