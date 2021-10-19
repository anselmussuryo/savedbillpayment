package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
)

var testQueries *Queries

const (
	dbDriver = "posgres"
	dbSource = "mysql://adminDB:simaS123@tcp(localhost:3307)/savedbillpaymentdb"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
