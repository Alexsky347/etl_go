package repositories

import (
	"context"
	"etl-go/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type ClientRepository struct {
	DB *pgxpool.Pool
}

func (r *ClientRepository) FindAll() ([]*models.Client, error) {
	if r.DB == nil {
		return nil, errors.New("database connection pool is closed")
	}
	rows, err := r.DB.Query(context.Background(), "SELECT * FROM clients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []*models.Client
	for rows.Next() {
		var client *models.Client
		err = rows.Scan(&client.ID, &client.Num1, &client.Num2, &client.Num3, &client.Subsidiary, &client.CreatedAt, &client.UpdatedAt)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return clients, nil
}