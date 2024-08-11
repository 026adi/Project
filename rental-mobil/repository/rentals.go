package repository

import (
	"database/sql"
	"rental-mobil/structs"
)

func GetAllRentals(db *sql.DB) ([]structs.Rental, error) {
	rows, err := db.Query("SELECT rental_id, user_id, car_id, rental_date, return_date, total_cost, status FROM rentals")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rentals []structs.Rental
	for rows.Next() {
		var rental structs.Rental
		if err := rows.Scan(&rental.RentalID, &rental.UserID, &rental.CarID, &rental.RentalDate, &rental.ReturnDate, &rental.TotalCost, &rental.Status); err != nil {
			return nil, err
		}
		rentals = append(rentals, rental)
	}

	return rentals, nil
}

func InsertRental(db *sql.DB, rental structs.Rental) error {
	query := "INSERT INTO rentals (user_id, car_id, rental_date, return_date, total_cost, status) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := db.Exec(query, rental.UserID, rental.CarID, rental.RentalDate, rental.ReturnDate, rental.TotalCost, rental.Status)
	return err
}

func UpdateRental(db *sql.DB, rental structs.Rental) error {
	query := "UPDATE rentals SET user_id = $1, car_id = $2, rental_date = $3, return_date = $4, total_cost = $5, status = $6 WHERE rental_id = $7"
	_, err := db.Exec(query, rental.UserID, rental.CarID, rental.RentalDate, rental.ReturnDate, rental.TotalCost, rental.Status, rental.RentalID)
	return err
}

func DeleteRental(db *sql.DB, id int) error {
	query := "DELETE FROM rentals WHERE rental_id = $1"
	_, err := db.Exec(query, id)
	return err
}
