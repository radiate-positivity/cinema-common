package models

type Movie struct {
	ID          int64    `json:"id"`
	Title       string   `json:"title"`
	AgeRating   string   `json:"ageRating"`
	PosterURL   string   `json:"posterUrl"`
	Description string   `json:"description"`
	Duration    int      `json:"duration"`
	Director    string   `json:"director"`
	Cast        []string `json:"cast"`
}