package dbconnector

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const dbname = "mysql"
const dburl = "root:Krackravi@15@tcp(localhost:3306)/temp"

func Dbconnector() {
	// Open a database connection
	db, err := sql.Open(dbname, dburl)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to the database!")
}
