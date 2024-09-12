package golangdatabasemysql

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExecSql(t *testing.T) {
	db, _ := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES('roj', 'Muhamad Sirojudin', 'siroj@hyperscal.com', 100000, 90, '2000-11-11', true), ('izz', 'Izz Luthfi El Shirazy', 'izz@luthfi.com', 3000000, 100, '2024-11-11', true)"
	_, err := db.ExecContext(ctx, query)

	if err != nil {
		assert.Fail(t, "Failed to insert data", err)
	}

	fmt.Println("Success insert new data")
}

func TestQuerySql(t *testing.T) {
	db, _ := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		assert.Fail(t, "Failed to get data", err)
	}

	fmt.Println("Success get data")

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt) // sesuai urutan select
		if err != nil {
			assert.Fail(t, err.Error())
		}
		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
		if email.Valid {
			fmt.Println("email: ", email.String)
		}
		fmt.Println("balance: ", balance)
		fmt.Println("rating: ", rating)
		if birthDate.Valid {
			fmt.Println("birthDate: ", birthDate.Time)
		}
		fmt.Println("married: ", married)
		fmt.Println("createdAt: ", createdAt)
		fmt.Println()
	}
	defer rows.Close()
}

func TestSqlInjection(t *testing.T) {
	db, _ := GetConnection()
	defer db.Close()

	ctx := context.Background()
	username := "admin'; #" //causing sql injection
	password := "salah"

	query := "SELECT username FROM user WHERE username = '" + username + "' AND password ='" + password + "' LIMIT 1"
	fmt.Println(query) //SELECT username FROM user WHERE username = 'admin'; #' AND password ='salah' LIMIT 1
	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		assert.Fail(t, "Failed to get data", err)
	}

	fmt.Println("Success get data")

	if rows.Next() {
		var username string

		err := rows.Scan(&username)
		if err != nil {
			assert.Fail(t, err.Error())
		}
		fmt.Println("Sukses login, with username: ", username)
	} else {
		fmt.Println("Gagal login")
	}
	defer rows.Close()
}

func TestQueryWithParameter(t *testing.T) {

	db, _ := GetConnection()
	defer db.Close()

	ctx := context.Background()
	username := "admin'; #" //causing sql injection
	password := "salah"

	query := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, query, username, password)

	if err != nil {
		assert.Fail(t, "Failed to get data", err)
	}

	fmt.Println("Success get data")

	if rows.Next() {
		var username string

		err := rows.Scan(&username)
		if err != nil {
			assert.Fail(t, err.Error())
		}
		fmt.Println("Sukses login, with username: ", username)
	} else {
		fmt.Println("Gagal login")
	}
	defer rows.Close()
}

func TestExecSqlParameter(t *testing.T) {
	db, _ := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "siroj'; DROP TABLE user; #"
	password := "roj"

	query := "INSERT INTO user(username, password) VALUES(?, ?)"
	_, err := db.ExecContext(ctx, query, username, password)

	if err != nil {
		assert.Fail(t, "Failed to insert data", err)
	}

	fmt.Println("Success insert new data")
}
