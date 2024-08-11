package repository

import (
	"database/sql"
	"rental-mobil/structs"
)

func GetAllPayments(db *sql.DB) ([]structs.Payment, error) {
	rows, err := db.Query("SELECT payment_id, user_id, amount, payment_date, rental_id FROM payments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []structs.Payment
	for rows.Next() {
		var payment structs.Payment
		if err := rows.Scan(&payment.PaymentID, &payment.PaymentID, &payment.Amount, &payment.PaymentDate, &payment.RentalID); err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	return payments, nil
}

func InsertPayment(db *sql.DB, payment structs.Payment) error {
	query := "INSERT INTO payments (user_id, amount, payment_date, rental_id) VALUES ($1, $2, $3, $4)"
	_, err := db.Exec(query, payment.PaymentID, payment.Amount, payment.PaymentDate, payment.RentalID)
	return err
}

func UpdatePayment(db *sql.DB, payment structs.Payment) error {
	query := "UPDATE payments SET user_id = $1, amount = $2, payment_date = $3, rental_id = $4 WHERE payment_id = $5"
	_, err := db.Exec(query, payment.PaymentID, payment.Amount, payment.PaymentDate, payment.RentalID, payment.PaymentID)
	return err
}

func DeletePayment(db *sql.DB, id int) error {
	query := "DELETE FROM payments WHERE payment_id = $1"
	_, err := db.Exec(query, id)
	return err
}
