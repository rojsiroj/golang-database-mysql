package golangdatabasemysql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
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

func TestAutoIncrement(t *testing.T) {
	db, _ := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "siroj@hyperscal"
	comment := "Test Comment"

	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, query, email, comment)

	if err != nil {
		assert.Fail(t, "Failed to insert data", err)
	}

	insertId, err := result.LastInsertId()

	if err != nil {
		assert.Fail(t, "Failed to insert data", err)
	}

	fmt.Println("Success insert new data with id: ", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db, _ := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	statement, err := db.PrepareContext(ctx, query) // prepare only use one connection, suitable for inserting large amounts of data
	if err != nil {
		assert.Fail(t, err.Error())
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "roj" + strconv.Itoa(i) + "gmail.com"
		comment := "Komentar ke-" + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)

		if err != nil {
			assert.Fail(t, err.Error())
		}

		id, err := result.LastInsertId()

		if err != nil {
			assert.Fail(t, err.Error())
		}

		fmt.Println("Success insert new comment data with id: ", id)
	}
}

func TestTransaction(t *testing.T) {
	db, _ := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		assert.Fail(t, err.Error())
	}

	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	for i := 0; i < 10; i++ {
		email := "roj" + strconv.Itoa(i) + "gmail.com"
		comment := "Komentar ke-" + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, query, email, comment)

		if err != nil {
			assert.Fail(t, err.Error())
		}

		id, err := result.LastInsertId()

		if err != nil {
			assert.Fail(t, err.Error())
		}

		fmt.Println("Success insert new comment data with id: ", id)
	}

	err = tx.Commit()
	// err = tx.Rollback()
	if err != nil {
		assert.Fail(t, err.Error())
	}
}
