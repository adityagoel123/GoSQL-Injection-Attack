package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	msg := string(data)

	fmt.Println("The data has been read from the Standard Input.", msg)

	db, err := sql.Open("postgres", "user=postgres password=s3cr3t sslmode=disable")

	if err != nil {
		fmt.Println("There is some error in opening DB connection.", err)
		log.Fatal(err)
	} else {
		fmt.Println("The connection to the database has been opened.", db)
	}
	defer db.Close()

	if err := createTables(db); err != nil {
		fmt.Println("There is some Error whie creating tables.", err)
		log.Fatal(err)
	}

	if err := insertLog(db, time.Now(), msg); err != nil {
		fmt.Println("There is some Error while Inserting Records into DB.", err)
		log.Fatal(err)
	}
}
