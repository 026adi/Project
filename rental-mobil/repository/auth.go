package repository

import (
	"database/sql"
	"rental-mobil/structs"
)

func Register(db *sql.DB, user structs.User) error {
	sql := "INSERT INTO Users (username, password, email) VALUES ($1, $2, $3)"
	_, err := db.Exec(sql, user.Username, user.Password, user.Email)
	return err
}

func GetUserByUsername(db *sql.DB, username string) (structs.User, error) {
	var user structs.User
	sql := "SELECT * FROM Users WHERE username = $1"
	row := db.QueryRow(sql, username)
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, err
}

func CreateUser(db *sql.DB, user structs.User) (structs.User, error) {
	query := `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`
	err := db.QueryRow(query, user.Username, user.Password).Scan(&user.ID)
	if err != nil {
		return user, err
	}
	return user, nil
}
