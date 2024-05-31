package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/ClickHouse/clickhouse-go"

	"statisticsservice/schemas"
)

var DB *sql.DB
var TableName = "events"

func ConnectToDb() {
	dsn := os.Getenv("DB_URL")
	var err error
	DB, err = sql.Open("clickhouse", dsn)
	if err != nil {
		log.Fatal(err)
	}

	schema := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s (
            task_id      Int64,
            username      String,
            event_type   Enum8('View' = 0, 'Like' = 1)
        ) ENGINE = ReplacingMergeTree()
        ORDER BY (task_id, username, event_type)
    `, TableName)

	_, err = DB.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}

func AddEvent(event schemas.Event) {
	tx, err := DB.Begin()
	if err != nil {
		log.Printf("Error adding event: %v", err)
	}
	defer tx.Rollback()

	_, err = tx.Exec(fmt.Sprintf("INSERT INTO %s (task_id, username, event_type) VALUES (?, ?, ?) FINAL", TableName),
		event.TaskID, event.Username, int8(event.EventType))
	if err != nil {
		log.Printf("Error adding event: %v", err)
	}

	if err = tx.Commit(); err != nil {
		log.Printf("Error adding event: %v", err)
	}
}
