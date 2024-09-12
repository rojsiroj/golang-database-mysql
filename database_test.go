package golangdatabasemysql

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseConnectMysql(t *testing.T) {
	db, err := GetConnection()

	if err != nil {
		assert.Fail(t, "Failed to Connect database", err)
	}
	defer db.Close()
}
