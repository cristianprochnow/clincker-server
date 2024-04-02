package models

import (
	"clincker/db"
	"fmt"
)

type UserStruct struct {
	Id      int    `json:"id"`
	Email   string `json:"email"`
	IsAdmin string `json:"is_admin"`
}

type UserModel struct {
	List func() ([]UserStruct, error)
}

func User() UserModel {
	return UserModel{
		List: list,
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

	for rows.Next() {
		var user UserStruct

		exception := rows.Scan(
			&user.Id, &user.Email, &user.IsAdmin,
		)

		if exception != nil {
			return nil, fmt.Errorf("models.users.list: %s", exception.Error())
		}

		users = append(users, user)
	}

	exception = rows.Err()

	if exception != nil {
		return nil, fmt.Errorf("models.users.list: %s", exception.Error())
	}

	return users, nil
}
