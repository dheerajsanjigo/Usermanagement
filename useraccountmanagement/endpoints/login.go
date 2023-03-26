package useraccountmanagement

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Check if request method is POST
	if r.Method != http.MethodPost {
		// If not, throw error
		http.Error(w, "Not Post Request", http.StatusBadRequest)
		return
	}

	// Decode JSON request body into LoginRequest struct
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
		return
	}

	// Check if user exists in database
	var storedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE Username = ?", req.Username).Scan(&storedPassword)
	if err == sql.ErrNoRows {
		http.Error(w, "Invalid Username ", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Compare stored password with input password
	if storedPassword != req.Password {
		http.Error(w, "Invalid  password", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Logged in successfully")
	w.WriteHeader(http.StatusOK)
}