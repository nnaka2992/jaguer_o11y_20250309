package sql

import (
	"fmt"
	"os"
	"database/sql"

	"github.com/XSAM/otelsql"
	_ "github.com/jackc/pgx/v5/stdlib"

)
var DB *Queries

func Init() error {
	db, err := Open()
	if err != nil {
		return err
	}

	DB = New(db)
	return nil
}

func Open() (*sql.DB ,error) {
	db, err := otelsql.Open(
		"pgx", os.Getenv("DATABASE_URL"), otelsql.WithSQLCommenter(true),
	)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v\n", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Unable to ping database: %v\n", err)
	}
	
	err = otelsql.RegisterDBStatsMetrics(
		db, otelsql.WithSQLCommenter(true),
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
