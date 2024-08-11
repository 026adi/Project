package structs

type Car struct {
	CarID              int    `json:"car_id"`
	Make               string `json:"make"`
	Model              string `json:"model"`
	Year               int    `json:"year"`
	RegistrationNumber string `json:"registration_number"`
	Available          bool   `json:"available"`
}
