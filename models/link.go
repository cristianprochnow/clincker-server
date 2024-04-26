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
}

type LinkModel struct {
	ListByUser func(userId int) ([]LinkStruct, error)
}

func Link() LinkModel {
	return LinkModel{
		ListByUser: linkByUser,
	}
}

func linkByUser(userId int) ([]LinkStruct, error) {
	var links []LinkStruct

	sql := db.Connect()
	rows, exception := sql.Query(
		"SELECT links.id, links.hash, links.created_at,"+
			"links.edited_at, links.original_url, links.domain,"+
			"links.resources, links.protocol"+
			"FROM links "+
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
		)

		if exception != nil {
			return nil, fmt.Errorf("models.links.list: %s", exception.Error())
		}

		links = append(links, link)
	}

	return links, nil
}

func addLinkUser(link LinkInsertStruct) (int, error) {
	sql := db.Connect()

	insertResult, exception := sql.ExecContext(
		context.Background(),
		"INSERT INTO users("+
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

func updateLinkUser(link LinkInsertStruct, id int) (int, error) {
	sql := db.Connect()

	insertResult, exception := sql.ExecContext(
		context.Background(),
		"UPDATE links "+
			"SET hash = ?, original_url = ?, domain = ?, "+
			"resources = ?, protocol = ?, user = ?, edited_at = NOW()"+
			"WHERE id = ?",
		link.Hash, link.OriginalUrl, link.Domain,
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
