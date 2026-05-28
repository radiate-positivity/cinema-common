package models

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	CardLast4 string `json:"cardLast4"`
	Role      string `json:"role"`
	CreatedAt string `json:"createdAt,omitempty"`
}