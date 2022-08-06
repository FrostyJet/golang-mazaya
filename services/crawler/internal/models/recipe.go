package models

type Recipe struct {
	Title               string          `json:"title"`
	PreparationTime     int             `json:"preparation_time"`
	PreparationTimeUnit string          `json:"preparation_time_unit"`
	CookingTime         int             `json:"cooking_time"`
	CookingTimeUnit     string          `json:"cooking_time_unit"`
	Serves              int             `json:"serves"`
	Difficulty          string          `json:"difficulty"`
	DescriptionText     string          `json:"description_text"`
	DescriptionHTML     string          `json:"description_html"`
	Poster              string          `json:"poster"`
	Keywords            []RecipeKeyword `json:"keywords"`
}

type RecipeKeyword string
