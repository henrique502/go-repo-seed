package postgre

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var instance *pgx.Conn

// Init database connection
func init() {
	host := os.Getenv("POSTGRE_HOST")
	port := os.Getenv("POSTGRE_PORT")
	user := os.Getenv("POSTGRE_USER")
	password := os.Getenv("POSTGRE_PASSWORD")
	dbname := os.Getenv("POSTGRE_DATABASE_NAME")

	var err error
	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, dbname)
	instance, err = pgx.Connect(context.Background(), connection)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer instance.Close(context.Background())

	err = instance.Ping(context.Background())
	if err != nil {
		fmt.Println("pq error:", err)
		panic(err)
	}
}
