package dbconnector

import (
	"database/sql"
	"fmt"
	types "github.com/dheerajsanjigo/Usermanagement/models"

	_ "github.com/go-sql-driver/mysql"
)

// Query for creating user
func CreateUser(db *sql.DB, user types.User) error {
	fmt.Println(user)
	// Create a prepared statement to insert a new user into the database
	stmt, err := db.Prepare("INSERT INTO users (Email, Password,Username,Fullname) VALUES (?, ?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the prepared statement to insert the new user into the database
	_, err = stmt.Exec(user.Email, user.Password, user.Username, user.FullName)
	if err != nil {
		return err
	}

	return nil
}

// Query for updating userdetails
func UpdateUser(db *sql.DB, user types.User) error {
	fmt.Println(user)
	// Create a prepared statement to update the  user password into the database
	stmt, err := db.Prepare("UPDATE `users` SET `Password` = ? WHERE `Username` = ?;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the prepared statement to update the  user password into the database
	_, err = stmt.Exec(user.NPassword, user.Username)
	if err != nil {
		return err
	}

	return nil

}

// Query for deleting user
func DeleteUser(db *sql.DB, user string) error {
	fmt.Println(user)
	// Create a prepared statement to update the  user password into the database
	stmt, err := db.Prepare("DELETE FROM users WHERE Username= ? ;")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the prepared statement to update the  user password into the database
	_, err = stmt.Exec(user)
	if err != nil {
		return err
	}

	return nil
}
