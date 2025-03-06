package main

import (
	"context"
	
	//	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

 	"go.opentelemetry.io/otel"

	"github.com/nnaka2992/jaguer_o11y_20250307/src/sql"
)



func main() {
	shutdown, err := initTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer shutdown()

	db, err := sql.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


	ctx := context.Background()

	slog.InfoContext(ctx, "Starting example")
	ctx, span := otel.Tracer(instrumentationName).Start(ctx, "parent")
	defer span.End()
	// run 5 times
	for i := 0; i < 5; i++ {
		if err = run(db, ctx); err != nil {
			slog.ErrorContext(ctx, "Example failed", slog.Any("error", err))
			os.Exit(1)
		}
	}
}

func run(db *sql.DB, ctx context.Context) error {
	ctx, span := otel.Tracer(instrumentationName).Start(ctx, "child")
	defer span.End()
	err := query(ctx, db)
	if err != nil {
			span.RecordError(err)
		return err
	}
	return nil
}

func query(ctx context.Context, db *sql.DB) error {
	// setup table named tab
//	_, errn := db.ExecContext(ctx, "create table if not exists tab (greeting text)")
//	if errn != nil {
//		return errn
//	}
//	_, errn = db.ExecContext(ctx, "insert into tab (greeting) values ('Hello, world!')")
//	if errn != nil {
//		return errn
//	}
	// Make a query

	rows, err := db.QueryContext(ctx, "select * from tab")
	if err != nil {
		return err
	}
	defer rows.Close()

 var greeting string
	for rows.Next() {
		err = rows.Scan(&greeting)
		if err != nil {
			return err
		}
	}
	fmt.Println(greeting)
	return nil
}
