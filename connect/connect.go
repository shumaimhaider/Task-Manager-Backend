package connect

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func Connect() *bun.DB {
	connection := pgdriver.NewConnector(
		pgdriver.WithAddr("localhost:5432"),
		pgdriver.WithUser("postgres"),
		pgdriver.WithPassword("postgres"),
		pgdriver.WithDatabase("Task-Manager"),
	)
	sqlDB := sql.OpenDB(connection)
	bunDB := bun.NewDB(sqlDB, pgdialect.New(), bun.WithDiscardUnknownColumns())
	debug := bundebug.NewQueryHook(bundebug.WithVerbose(true))
	bunDB.AddQueryHook(debug)
	return bunDB
}
