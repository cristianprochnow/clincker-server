package models

import (
	"clincker/db"
	"context"
	"fmt"
)

type LinkInsertStruct struct {
	Hash        string `json:"hash"`
	OriginalUrl string `json:"original_url"`
	Domain      string `json:"domain"`
	Resources   string `json:"resources"`
	Protocol    string `json:"protocol"`
	User        int    `json:"user"`
}

type LinkStruct struct {
	Id          int    `json:"id"`
	Hash        string `json:"hash"`
	CreatedAt   string `json:"created_at"`
	EditedAt    string `json:"edited_at"`
	OriginalUrl string `json:"original_url"`
	Domain      string `json:"domain"`
	Resources   string `json:"resources"`
	Protocol    string `json:"protocol"`
	UserId      int    `json:"user_id"`
	UserName    string `json:"user_name"`
}

type NewLinkResponseStruct struct {
	Id   int    `json:"id"`
	Hash string `json:"hash"`
}

type LinkModel struct {
	ListByUser func(userId int) ([]LinkStruct, error)
	Create     func(link LinkInsertStruct) (int, error)
	Update     func(link LinkInsertStruct, id int) (int, error)
	Show       func(id int) (*LinkStruct, error)
	Delete     func(id int) (bool, error)
	IsValid    func(dataSent LinkInsertStruct) bool
}

func Link() LinkModel {
	return LinkModel{
		ListByUser: linkByUser,
		Create:     createLink,
		Update:     updateLink,
		Show:       showLink,
		Delete:     deleteLink,
		IsValid:    isValidLink,
	}
}

func linkByUser(userId int) ([]LinkStruct, error) {
	var links []LinkStruct

	sql := db.Connect()
	rows, exception := sql.Query(
		"SELECT links.id, links.hash, links.created_at, "+
			"COALESCE(links.edited_at, ''), links.original_url, links.domain, "+
			"links.resources, links.protocol, links.user user_id, "+
			"users.name user_name "+
			"FROM links "+
			"LEFT JOIN users ON users.id = links.user "+
			"WHERE links.user = ?", userId)

	if exception != nil {
		return nil, fmt.Errorf("models.links.list: %s", exception)
	}

	defer rows.Close()

	exception = rows.Err()

	if exception != nil {
		return nil, fmt.Errorf("models.links.list: %s", exception.Error())
	}

	for rows.Next() {
		var link LinkStruct

		exception := rows.Scan(
			&link.Id,
			&link.Hash,
			&link.CreatedAt,
			&link.EditedAt,
			&link.OriginalUrl,
			&link.Domain,
			&link.Resources,
			&link.Protocol,
			&link.UserId,
			&link.UserName,
		)

		if exception != nil {
			return nil, fmt.Errorf("models.links.list: %s", exception.Error())
		}

		links = append(links, link)
	}

	return links, nil
}

func createLink(link LinkInsertStruct) (int, error) {
	sql := db.Connect()

	insertResult, exception := sql.ExecContext(
		context.Background(),
		"INSERT INTO links("+
			"hash, original_url, domain, resources, protocol, user"+
			") VALUES (?, ?, ?, ?, ?, ?)",
		link.Hash, link.OriginalUrl, link.Domain,
		link.Resources, link.Protocol, link.User)

	if exception != nil {
		return 0, fmt.Errorf("models.links.create: %s", exception.Error())
	}

	id, exception := insertResult.LastInsertId()

	if exception != nil {
		return 0, fmt.Errorf("models.links.create: %s", exception.Error())
	}

	return int(id), nil
}

func updateLink(link LinkInsertStruct, id int) (int, error) {
	sql := db.Connect()

	insertResult, exception := sql.ExecContext(
		context.Background(),
		"UPDATE links "+
			"SET original_url = ?, domain = ?, "+
			"resources = ?, protocol = ?, user = ?, edited_at = NOW()"+
			"WHERE id = ?",
		link.OriginalUrl, link.Domain,
		link.Resources, link.Protocol, link.User,
		id)

	if exception != nil {
		return 0, fmt.Errorf("models.users.update: %s", exception.Error())
	}

	idExists, exception := insertResult.LastInsertId()

	if exception != nil {
		return 0, fmt.Errorf("models.users.update: %s", exception.Error())
	}

	return int(idExists), nil
}

func showLink(id int) (*LinkStruct, error) {
	var link LinkStruct

	sql := db.Connect()
	exception := sql.QueryRow(
		"SELECT links.id, links.hash, links.created_at, "+
			"COALESCE(links.edited_at, ''), links.original_url, links.domain, "+
			"links.resources, links.protocol, links.user user_id, "+
			"users.name user_name "+
			"FROM links "+
			"LEFT JOIN users ON users.id = links.user "+
			"WHERE links.id = ?", id,
	).Scan(
		&link.Id,
		&link.Hash,
		&link.CreatedAt,
		&link.EditedAt,
		&link.OriginalUrl,
		&link.Domain,
		&link.Resources,
		&link.Protocol,
		&link.UserId,
		&link.UserName,
	)

	if exception != nil {
		return nil, fmt.Errorf("models.links.show: %s", exception.Error())
	}

	return &link, nil
}

func deleteLink(id int) (bool, error) {
	success := false
	sql := db.Connect()
	deleteResult, exception := sql.ExecContext(
		context.Background(),
		"DELETE FROM links WHERE id = ?", id)

	if exception != nil {
		return success, fmt.Errorf("models.users.delete: %s", exception.Error())
	}

	rowsAffected, _ := deleteResult.RowsAffected()

	if rowsAffected > 0 {
		success = true
	}

	return success, nil
}

func isValidLink(dataSent LinkInsertStruct) bool {
	return dataSent.User != 0 &&
		dataSent.OriginalUrl != "" &&
		dataSent.Domain != "" &&
		dataSent.Resources != "" &&
		dataSent.Protocol != ""
}
