package repository

import (
	"database/sql"
	"golang-mazaya/storefront/internal/domain"
	"strings"
	"sync"
	"time"
)

var once sync.Once

type recipeRepository struct {
	db *sql.DB
}

var instance domain.RecipeRepository

func NewRecipeRepository(db *sql.DB) domain.RecipeRepository {
	once.Do(func() {
		instance = &recipeRepository{
			db: db,
		}
	})

	return instance
}

func (r *recipeRepository) Create(data *domain.Recipe) (*domain.Recipe, error) {
	query := `INSERT INTO recipe(
		title, preparation_time, preparation_time_unit, cooking_time, cooking_time_unit,
		serves, difficulty, description_text, description_html, poster, date_created
	) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	err = r.Validate(data)
	if err != nil {
		return nil, err
	}

	lastInsertId := 0
	err = stmt.QueryRow(
		data.Title, data.PreparationTime, data.PreparationTimeUnit, data.CookingTime, data.CookingTimeUnit,
		data.Serves, data.Difficulty, data.DescriptionText, data.DescriptionHTML, data.Poster, time.Now(),
	).Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}

	data.ID = lastInsertId

	return data, nil
}

func (r *recipeRepository) Validate(data *domain.Recipe) error {
	data.DescriptionText = strings.TrimSpace(data.DescriptionText)
	data.DescriptionHTML = strings.TrimSpace(data.DescriptionHTML)

	return nil
}
