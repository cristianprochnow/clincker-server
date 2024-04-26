package models

import (
	"clincker/db"
	"fmt"
)

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
