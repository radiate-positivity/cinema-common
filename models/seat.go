package models

type Seat struct {
	ID          int  `json:"id"`
	Row         int  `json:"row"`
	Number      int  `json:"number"`
	HallID      int  `json:"hallId"`
	IsAvailable bool `json:"isAvailable"`
}