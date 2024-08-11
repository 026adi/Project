package structs

type Rental struct {
	RentalID   int     `json:"rental_id"`
	UserID     int     `json:"user_id"`
	CarID      int     `json:"car_id"`
	RentalDate string  `json:"rental_date"`
	ReturnDate string  `json:"return_date"`
	TotalCost  float64 `json:"total_cost"`
	Status     string  `json:"status"`
}
