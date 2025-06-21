package repository

import (
	"github.com/gabrielmartins0312/crud-go-with-api-and-db/config"
	"github.com/gabrielmartins0312/crud-go-with-api-and-db/model"
)

func CreateUser(user model.User) error {
	_, err := config.DB.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	return err
}

func GetAllUsers() ([]model.User, error) {
	rows, err := config.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		rows.Scan(&u.ID, &u.Name, &u.Email)
		users = append(users, u)
	}
	return users, nil
}

func UpdateUser(user model.User) error {
	_, err := config.DB.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, user.ID)
	return err
}
