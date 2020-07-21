package models

// Donation represent donation data
type Donation struct {
	ID          int    `json:"id"`
	IDYayasan   int    `json:"id_yayasan"`
	Amount      int    `json:"amount"`
	PhoneNumber string `json:"phone_number"`
	Category    string `json:"category"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}
