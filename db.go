package main

import (
	"database/sql"
	_ "embed"
	"fmt"
	"time"
)

var (
	//go:embed schema.sql
	schemaSQL string

	//go:embed insert.sql
	insertSQL string
)

func createTables(db *sql.DB) error {
	fmt.Println("The schemaSQL thus read is : ", schemaSQL)
	_, err := db.Exec(schemaSQL)
	return err
}

func insertLog(db *sql.DB, time time.Time, message string) error {
	fmt.Println("The insertSQL thus read from FILE is : ", insertSQL)
	_, err := db.Exec(insertSQL, time, message)
	return err
}
