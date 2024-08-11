package repository

import (
	"database/sql"
	"rental-mobil/structs"
)

func GetAllCars(db *sql.DB) ([]structs.Car, error) {
	var cars []structs.Car
	sql := "SELECT * FROM Cars"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var car structs.Car
		err := rows.Scan(&car.CarID, &car.Make, &car.Model, &car.Year, &car.RegistrationNumber, &car.Available)
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	return cars, nil
}

func InsertCar(db *sql.DB, car structs.Car) error {
	sql := "INSERT INTO Cars (make, model, year, registration_number, available) VALUES ($1, $2, $3, $4, $5)"
	_, err := db.Exec(sql, car.Make, car.Model, car.Year, car.RegistrationNumber, car.Available)
	return err
}

func UpdateCar(db *sql.DB, car structs.Car) error {
	sql := "UPDATE Cars SET make = $1, model = $2, year = $3, registration_number = $4, available = $5 WHERE car_id = $6"
	_, err := db.Exec(sql, car.Make, car.Model, car.Year, car.RegistrationNumber, car.Available, car.CarID)
	return err
}

func DeleteCar(db *sql.DB, carID int) error {
	sql := "DELETE FROM Cars WHERE car_id = $1"
	_, err := db.Exec(sql, carID)
	return err
}
