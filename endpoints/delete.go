package endpoints

import (
	"database/sql"
	"fmt"
	"net/http"

	dbconnector "github.com/dheerajsanjigo/Usermanagement/dbconnector"
	_ "github.com/go-sql-driver/mysql"
)

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := dbconnector.Dbconnector()
	defer db.Close()
	// Check if request method is DELETE
	if r.Method != http.MethodDelete {
		// If not, throw error
		http.Error(w, "Not Delete Request", http.StatusBadRequest)
		return
	}

	userName := r.URL.Query().Get("username")

	// Check if user exists in database
	var storedPassword string
	err = db.QueryRow("SELECT password FROM users WHERE Username = ?", userName).Scan(&storedPassword)
	if err == sql.ErrNoRows {
		http.Error(w, "Invalid Username ", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := dbconnector.DeleteUser(db, userName); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Deleted successfully")
	w.WriteHeader(http.StatusOK)

}
