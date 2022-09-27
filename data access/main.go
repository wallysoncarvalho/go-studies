package main

import (
	"database/sql"
	//"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.TraceLevel)

	if _, logJson := os.LookupEnv("GO_LOG_LEVEL"); logJson {
		log.SetFormatter(&log.JSONFormatter{})
	}
	

	log.Info("Starting database connection!")


	logContext := log.WithFields(log.Fields{"field1": "this is field 1"})
	
	logContext.Info(" teste logging")

	logContext.Warning("Careful what you wish for!")

	db, err := sql.Open("postgres", "postgresql://postgres:postgres@localhost:54322/meu_banco?sslmode=disable")
	
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("error pinging to database: %v", err)
	}

	rows, err := db.Query("SELECT * FROM entry;")

	var entrie Entry

	for rows.Next() {
		rows.Scan(&entrie.id, &entrie.account, &entrie.amount, &entrie.createdAt)
		// fmt.Println(entrie)

		logContext.Info(entrie)
	}



}

type Entry struct {
	id int
	account uint64
	amount int64
	createdAt time.Time
}
