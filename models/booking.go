package models

type Booking struct {
	ID         int64  `json:"id"`
	UserID     int64  `json:"userId"`
	ShowtimeID int64  `json:"showtimeId"`
	Seats      string `json:"seats"`
	Status     string `json:"status"`
	TotalPrice int    `json:"totalPrice"`
	CreatedAt  string `json:"createdAt"`
}