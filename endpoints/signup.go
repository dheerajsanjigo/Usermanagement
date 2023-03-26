package endpoints

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	dbconnector "github.com/dheerajsanjigo/Usermanagement/dbconnector"
	validators "github.com/dheerajsanjigo/Usermanagement/validators"

	types "github.com/dheerajsanjigo/Usermanagement/models"
	_ "github.com/go-sql-driver/mysql"
)

var err error

func createUser(db *sql.DB, user types.User) error {
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

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	// Check if request method is POST
	if r.Method != http.MethodPost {
		// If not, throw error
		http.Error(w, "Not Post Request", http.StatusBadRequest)
		return
	}
	db, err := dbconnector.Dbconnector()
	defer db.Close()
	// Decode JSON request body into User struct
	var req types.User
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
		return
	}
	fmt.Println(req)
	if len(req.Password) < 8 {
		http.Error(w, "password must be at least 8 characters long", http.StatusBadRequest)
		return
	}

	val := validators.ValidateEmail(req.Email)
	if val == false {
		http.Error(w, "Invalid email Format", http.StatusBadRequest)
		return

	}
	if req.Username == "" {
		http.Error(w, "Email cannot be empty", http.StatusBadRequest)
		return

	}

	// Call the createUser function to insert the new user into the database
	if err := createUser(db, req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "created successfully")
	w.WriteHeader(http.StatusCreated)

}
