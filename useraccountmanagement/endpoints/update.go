package useraccountmanagement

import (
	
	"encoding/json"
	"fmt"
	"net/http"
	""



)

func UpdateUser(user UpdateRequest) error {
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

func updatePasswordHandler(w http.ResponseWriter, r *http.Request) {
	// Check if request method is PUT
	if r.Method != http.MethodPut {
		// If not, throw error
		http.Error(w, "Not Put Request", http.StatusBadRequest)
		return
	}

	// Decode JSON request body into LoginRequest struct
	var req UpdateRequest
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

	if req.NPassword != req.CNPassword {
		http.Error(w, "passwords didn't match", http.StatusBadRequest)
		return
	}

	//calling the Update function
	if err := UpdateUser(req); err != nil {
		http.Error(w, "Invalid  password", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Updated Password successfully")
	w.WriteHeader(http.StatusOK)

}