package models

type Ticket struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"userId"`
	BookingID   int64  `json:"bookingId"`
	MovieTitle  string `json:"movieTitle"`
	HallName    string `json:"hallName"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	Seats       string `json:"seats"`
	QRCode      string `json:"qrCode"`
	IsUsed      bool   `json:"isUsed"`
	PurchasedAt string `json:"purchasedAt"`
}