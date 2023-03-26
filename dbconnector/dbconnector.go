package dbconnector

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const dbname = "mysql"
const dburl = "root:Krackravi@15@tcp(localhost:3306)/temp"

func Dbconnector() (*sql.DB, error) {
	// Open a database connection
	db, err := sql.Open(dbname, dburl)
	if err != nil {
		panic(err.Error())
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	return db, nil

}
