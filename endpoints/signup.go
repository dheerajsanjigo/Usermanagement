package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"

	dbconnector "github.com/dheerajsanjigo/Usermanagement/dbconnector"
	validators "github.com/dheerajsanjigo/Usermanagement/validators"

	types "github.com/dheerajsanjigo/Usermanagement/models"
	_ "github.com/go-sql-driver/mysql"
)

var err error

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

	// Validate password
	if err = validators.ValidatePassword(req.Password); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	val := validators.ValidateEmail(req.Email)
	if val == false {
		http.Error(w, "Invalid email Format", http.StatusBadRequest)
		return

	}
	if req.Username == "" {
		http.Error(w, "Username cannot be empty", http.StatusBadRequest)
		return

	}

	// Call the createUser function to insert the new user into the database
	if err := dbconnector.CreateUser(db, req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "created successfully")
	w.WriteHeader(http.StatusCreated)

}
