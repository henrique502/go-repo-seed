package database

import (
	"context"
	"fmt"
	"os"

	"github.com/henrique502/go-repo-seed/domain/alert"
	"github.com/henrique502/go-repo-seed/domain/integration"
	"github.com/henrique502/go-repo-seed/domain/team"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type DatabaseClient struct {
	client *pgx.Conn
}

func New(url string) DatabaseClient {
	connection, err := pgx.Connect(context.Background(), url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	err = connection.Ping(context.Background())
	if err != nil {
		fmt.Println("pq error:", err)
		panic(err)
	}

	return DatabaseClient{
		client: connection,
	}
}

func (db *DatabaseClient) Close() {
	db.client.Close(context.Background())
}

func (db *DatabaseClient) GetInstance() *DatabaseClient {
	return db
}

type Database interface {
	GetInstance() *DatabaseClient
	// Close connection
	Close()
	// AlertUpSert insert or update alert
	AlertUpSert(alert alert.Alert) (pgconn.CommandTag, error)
	// IntegrationUpSert insert or update integration
	IntegrationUpSert(integration integration.Integration) (pgconn.CommandTag, error)
	// TeamUpSert insert or update team
	TeamUpSert(team team.Team) (pgconn.CommandTag, error)
}
