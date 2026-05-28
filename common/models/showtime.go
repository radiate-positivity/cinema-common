package models

type Showtime struct {
	ID      int64  `json:"id"`
	MovieID int64  `json:"movieId"`
	Date    string `json:"date"`
	Time    string `json:"time"`
	Price   int    `json:"price"`
	HallID  int    `json:"hallId"`
}