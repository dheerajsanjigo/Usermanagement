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
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UpdateRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	NPassword  string `json:"NPassword"`
	CNPassword string `json:"CNPassword"`
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

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
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

func main() {
	//Create  the mysql connection
	db, err = sql.Open("mysql", "root:Krackravi@15@tcp(localhost:3306)/temp")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfully connected to DB")
	defer db.Close()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("called")
	})
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/UpdatePassword", updatePasswordHandler)
	http.HandleFunc("/DeleteUser", deleteUserHandler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

}
