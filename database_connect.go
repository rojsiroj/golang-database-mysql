package golangdatabasemysql

import (
	_ "github.com/go-sql-driver/mysql"

	"database/sql"
	"time"
)

func GetConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/learning_golang?parseTime=true")

	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db, nil
}
