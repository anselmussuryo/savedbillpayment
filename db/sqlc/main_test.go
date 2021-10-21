package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver = "mysql"
	dbSource = "adminDB:simaS123@(localhost:3307)/savedbillpaymentdb"
)

func TestMain(m *testing.M) {
	var err error

	testDB, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
