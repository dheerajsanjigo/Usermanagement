package useraccountmanagement

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
	Username string `json:"username"`
}

var db *sql.DB
var err error

func createUser(user SignupRequest) error {
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

func signupHandler(w http.ResponseWriter, r *http.Request) {
	// Check if request method is POST
	if r.Method != http.MethodPost {
		// If not, throw error
		http.Error(w, "Not Post Request", http.StatusBadRequest)
		return
	}

	// Decode JSON request body into SignupRequest struct
	var req SignupRequest
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
		return
	}
	fmt.Println(req)
	if len(req.Password) < 8 {
		http.Error(w, "password must be at least 8 characters long", http.StatusBadRequest)
		return
	}

	if req.Email == "" {
		http.Error(w, "Email cannot be empty", http.StatusBadRequest)
		return

	}
	if req.Username == "" {
		http.Error(w, "Email cannot be empty", http.StatusBadRequest)
		return

	}

	// Call the createUser function to insert the new user into the database
	if err := createUser(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "created successfully")
	w.WriteHeader(http.StatusCreated)

}
