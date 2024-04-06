package models

import (
	"clincker/db"
	"fmt"
)

type UserStruct struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	IsAdmin   string `json:"is_admin"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

type UserModel struct {
	List func() ([]UserStruct, error)
	Show func(id int) (*UserStruct, error)
}

func User() UserModel {
	return UserModel{
		List: list,
		Show: show,
	}
}

func list() ([]UserStruct, error) {
	var users []UserStruct

	sql := db.Connect()
	rows, exception := sql.Query("SELECT * FROM users")

	if exception != nil {
		return nil, fmt.Errorf("models.users.list: %s", exception)
	}

	defer rows.Close()

	exception = rows.Err()

	if exception != nil {
		return nil, fmt.Errorf("models.users.list: %s", exception.Error())
	}

	for rows.Next() {
		var user UserStruct

		exception := rows.Scan(
			&user.Id,
			&user.Email,
			&user.Name,
			&user.IsAdmin,
			&user.Password,
			&user.CreatedAt,
		)

		if exception != nil {
			return nil, fmt.Errorf("models.users.list: %s", exception.Error())
		}

		users = append(users, user)
	}

	return users, nil
}

func show(id int) (*UserStruct, error) {
	var user UserStruct

	sql := db.Connect()
	exception := sql.QueryRow(
		"SELECT * FROM users WHERE users.id = ?", id,
	).Scan(
		&user.Id,
		&user.Email,
		&user.Name,
		&user.IsAdmin,
		&user.Password,
		&user.CreatedAt,
	)

	if exception != nil {
		return nil, fmt.Errorf("models.users.show: %s", exception)
	}

	return &user, nil
}
