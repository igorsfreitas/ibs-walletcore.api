package database

import (
	"database/sql"

	"github.com/igorsfreitas/ibs-walletcore.api/internal/entity"
)

type ClientDB struct {
	DB *sql.DB
}

func NewClientDB(db *sql.DB) *ClientDB {
	return &ClientDB{DB: db}
}

func (c *ClientDB) GetClient(id string) (*entity.Client, error) {
	client := &entity.Client{}
	stmt, err := c.DB.Prepare(`SELECT id, name, email, created_at FROM clients WHERE id = $1`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	if err := row.Scan(&client.ID, &client.Name, &client.Email, &client.CreatedAt); err != nil {
		return nil, err
	}

	return client, nil
}

func (c *ClientDB) SaveClient(client *entity.Client) error {
	stmt, err := c.DB.Prepare(`INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, ?) RETURNING id`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if err := stmt.QueryRow(client.ID, client.Name, client.Email, client.CreatedAt).Scan(&client.ID); err != nil {
		return err
	}

	return nil
}
