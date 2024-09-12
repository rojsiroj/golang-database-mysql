package golangdatabasemysql

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecSql(t *testing.T) {
	db, _ := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer(id, name) VALUES('roj', 'Muhamad Sirojudin')"
	_, err := db.ExecContext(ctx, query)

	if err != nil {
		assert.Fail(t, "Failed to insert data", err)
	}

	fmt.Println("Success insert new data")
}
