package domain

type Recipe struct {
	ID                  int       `json:"id"`
	Title               string    `json:"title"`
	PreparationTime     int       `json:"preparation_time"`
	PreparationTimeUnit string    `json:"preparation_time_unit"`
	CookingTime         int       `json:"cooking_time"`
	CookingTimeUnit     string    `json:"cooking_time_unit"`
	Serves              int       `json:"serves"`
	Difficulty          string    `json:"difficulty"`
	DescriptionText     string    `json:"description_text"`
	DescriptionHTML     string    `json:"description_html"`
	Poster              string    `json:"poster"`
	Keywords            []Keyword `json:"keywords"`
}

type RecipeService interface {
	Create(r *Recipe) (*Recipe, error)
}

type RecipeRepository interface {
	Create(r *Recipe) (*Recipe, error)
}
