package models

import (
	"clincker/db"
	"context"
	"fmt"
)

type UserInsertStruct struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type NewUserResponseStruct struct {
	Id int `json:"id"`
}

type UserStruct struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	IsAdmin   string `json:"is_admin"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

type UserModel struct {
	List         func() ([]UserStruct, error)
	Show         func(id int) (*UserStruct, error)
	Verify       func(email string) (*UserStruct, error)
	Create       func(user UserInsertStruct) (int, error)
	Update       func(user UserInsertStruct, id int) (int, error)
	IsValid      func(dataSent UserInsertStruct) bool
	IsValidLogin func(dataSent UserLogin) bool
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginToken struct {
	Id    int    `json:"id"`
	Token string `json:"token"`
}

func User() UserModel {
	return UserModel{
		List:         list,
		Show:         show,
		Create:       create,
		Update:       update,
		Verify:       verify,
		IsValid:      isValidUser,
		IsValidLogin: isValidLogin,
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
		return nil, fmt.Errorf("models.users.show: %s", exception.Error())
	}

	return &user, nil
}

func verify(email string) (*UserStruct, error) {
	var user UserStruct

	sql := db.Connect()
	exception := sql.QueryRow(
		"SELECT * FROM users WHERE users.email = ?", email,
	).Scan(
		&user.Id,
		&user.Email,
		&user.Name,
		&user.IsAdmin,
		&user.Password,
		&user.CreatedAt,
	)

	if exception != nil {
		return nil, fmt.Errorf("models.users.show: %s", exception.Error())
	}

	if user.Id != 0 {
		return &user, fmt.Errorf(
			"models.users.verify: Conta com e-mail %s já existente",
			email,
		)
	}

	return nil, nil
}

func create(user UserInsertStruct) (int, error) {
	sql := db.Connect()

	insertResult, exception := sql.ExecContext(
		context.Background(),
		"INSERT INTO users(email, name, password) VALUES (?, ?, ?)",
		user.Email, user.Name, user.Password)

	if exception != nil {
		return 0, fmt.Errorf("models.users.create: %s", exception.Error())
	}

	id, exception := insertResult.LastInsertId()

	if exception != nil {
		return 0, fmt.Errorf("models.users.create: %s", exception.Error())
	}

	return int(id), nil
}

func update(user UserInsertStruct, id int) (int, error) {
	sql := db.Connect()

	insertResult, exception := sql.ExecContext(
		context.Background(),
		"UPDATE users SET email = ?, name = ?, password = ? WHERE id = ?",
		user.Email, user.Name, user.Password, id)

	if exception != nil {
		return 0, fmt.Errorf("models.users.update: %s", exception.Error())
	}

	idExists, exception := insertResult.LastInsertId()

	if exception != nil {
		return 0, fmt.Errorf("models.users.update: %s", exception.Error())
	}

	return int(idExists), nil
}

func isValidUser(dataSent UserInsertStruct) bool {
	return dataSent.Email != "" &&
		dataSent.Name != "" &&
		dataSent.Password != ""
}

func isValidLogin(dataSent UserLogin) bool {
	return dataSent.Email != "" && dataSent.Password != ""
}
