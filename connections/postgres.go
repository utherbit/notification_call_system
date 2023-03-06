package connections

import (
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"notificator/connections/appcontext"
	"notificator/utilities"
	"os"
	"time"
)

var (
	PostgresConn  *pgxpool.Pool
	PostgresConn2 *pgxpool.Pool
	PostgresErr   error
)

func InitPostgresConnection() {
	host, err := os.LookupEnv("POSTGRES_HOST")
	if !err {
		utilities.LogStrErr("POSTGRES-PANIC", "POSTGRES_HOST not found in .env file")
	}

	port, err := os.LookupEnv("POSTGRES_PORT")
	if !err {
		utilities.LogStrErr("POSTGRES-PANIC", "POSTGRES_PORT not found in .env file")
	}

	username, err := os.LookupEnv("POSTGRES_USERNAME")
	if !err {
		utilities.LogStrErr("POSTGRES-PANIC", "POSTGRES_USERNAME not found in .env file")
	}

	password, err := os.LookupEnv("POSTGRES_PASSWORD")
	if !err {
		utilities.LogStrErr("POSTGRES-PANIC", "POSTGRES_PASSWORD not found in .env file")
	}

	database, err := os.LookupEnv("POSTGRES_DATABASE_NAME")
	if !err {
		utilities.LogStrErr("POSTGRES-PANIC", "POSTGRES_DATABASE_NAME not found in .env file")
	}

	PostgresConn, PostgresErr = pgxpool.Connect(appcontext.Ctx, fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, database))
	utilities.PanicIfErr(PostgresErr)
	time.Sleep(1000)

	PostgresConn2, PostgresErr = pgxpool.Connect(appcontext.Ctx, fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, database))
	utilities.PanicIfErr(PostgresErr)

	utilities.LogStrInfo("Connect to Postgres was successfully")
}
