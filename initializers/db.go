package initializers

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	dsn := os.Getenv("DB_URL")

	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Database Connected")
	}
	// See "Important settings" section.
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
}
