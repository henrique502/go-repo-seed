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
	var err error
	instance, err = pgx.Connect(context.Background(), os.Getenv("POSTGRE_URL"))
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
