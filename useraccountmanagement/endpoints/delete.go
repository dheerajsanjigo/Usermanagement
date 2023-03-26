package endpoints

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func DeleteUser(user string) error {
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

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Check if request method is DELETE
	if r.Method != http.MethodDelete {
		// If not, throw error
		http.Error(w, "Not Delete Request", http.StatusBadRequest)
		return
	}

	userName := r.URL.Query().Get("username")

	// Check if user exists in database
	var storedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE Username = ?", userName).Scan(&storedPassword)
	if err == sql.ErrNoRows {
		http.Error(w, "Invalid Username ", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := DeleteUser(userName); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Deleted successfully")
	w.WriteHeader(http.StatusOK)

}
