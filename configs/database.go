package configs

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

var db *sql.DB

func ConnectToDB() *sql.DB {
	pqUrl, err := pq.ParseURL(os.Getenv("ELEPHANT_SQL"))

	if err != nil {
		log.Fatal(err)
	}

	db, err = sql.Open("postgres", pqUrl)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
