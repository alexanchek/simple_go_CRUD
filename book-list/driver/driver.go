package driver

import (
	"database/sql"
	"log"
	"os"
	"fmt"

	"github.com/lib/pq"
)

var db *sql.DB

func logFatal(err error) {
	if err !=nil {
		log.Fatal(err)
	}
}

// ConnectDB connect Database
func ConnectDB() *sql.DB {
	pgURL, err:= pq.ParseURL(os.Getenv("POSTRESSQL_URL"))
	log.Fatal(err)

	db, err = sql.Open("postgres", pgURL)
	if err!=nil {
		fmt.Println("An issue with opening")
		log.Fatal(err)
	}

	err = db.Ping()
	if err!=nil {
		fmt.Println("An issue with a ping with database")
		logFatal(err)
		log.Println(pgURL)
	}

	return db
}