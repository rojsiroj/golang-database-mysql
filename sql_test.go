package golangdatabasemysql

import (
	"context"
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
		var id, name, email string
		var balance int32
		var rating float64
		var birthDate, createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt) // sesuai urutan select
		if err != nil {
			assert.Fail(t, err.Error())
		}
		fmt.Println("Id: ", id)
		fmt.Println("Name: ", name)
		fmt.Println("email: ", email)
		fmt.Println("balance: ", balance)
		fmt.Println("rating: ", rating)
		fmt.Println("birthDate: ", birthDate)
		fmt.Println("married: ", married)
		fmt.Println("createdAt: ", createdAt)
	}

	defer rows.Close()
}
