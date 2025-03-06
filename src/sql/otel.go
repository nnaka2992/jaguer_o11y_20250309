package sql

import (
	"fmt"
	"database/sql"

	"github.com/XSAM/otelsql"
	_ "github.com/jackc/pgx/v5/stdlib"

)

func Open() (*sql.DB ,error) {
	db, err := otelsql.Open(
		"pgx", os.Getenv("DATABASE_URL"), otelsql.WithSQLCommenter(true),
	)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v\n", err)
	}

	err = otelsql.RegisterDBStatsMetrics(
		db, otelsql.WithSQLCommenter(true),
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
