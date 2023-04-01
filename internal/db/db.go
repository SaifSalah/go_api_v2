package db

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	Client *sqlx.DB
}

func NewDatabase() (*Database, error) {
	connectionStr := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		"localhost",
		"5432",
		"postgres",
		"uat_testing",
		"123123",
		"disable",
	)

	dbCconn, err := sqlx.Connect("postgres", connectionStr)
	if err != nil {
		return &Database{}, fmt.Errorf("cloud not connect to the databse %w", err)
	}
	return &Database{
		Client: dbCconn,
	}, nil
}

func (d *Database) Ping(ctx context.Context) error {
	return d.Client.PingContext(ctx)
}
