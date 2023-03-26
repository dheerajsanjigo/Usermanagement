package main

import (
	"fmt"
	"net/http"

	dbconnection "github.com/dheerajsanjigo/Usermanagement/dbconnector"
	endpoints "github.com/dheerajsanjigo/Usermanagement/useraccountmanagement/endpoints"
)

var err error

func main() {
	//Create  the mysql connection
	dbconnection.Dbconnector()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("called")
	})
	http.HandleFunc("/signup", endpoints.SignupHandler)
	http.HandleFunc("/login", endpoints.LoginHandler)
	http.HandleFunc("/UpdatePassword", endpoints.UpdatePasswordHandler)
	http.HandleFunc("/DeleteUser", endpoints.DeleteUserHandler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

}
