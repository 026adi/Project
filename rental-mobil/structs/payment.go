package structs

type Payment struct {
	PaymentID   int     `json:"payment_id"`
	RentalID    int     `json:"rental_id"`
	Amount      float64 `json:"amount"`
	PaymentDate string  `json:"payment_date"`
}
