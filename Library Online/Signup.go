package main

import (
	"fmt"
	"net/http"
)

type Account struct {
	Username string
	Password string
}

var accounts = make(map[string]Account)

func main() {
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil)
}

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	username := r.FormValue("name")
	password := r.FormValue("password")

	if _, exists := accounts[username]; exists {
		fmt.Fprintf(w, "Username already exists")
		return
	}

	accounts[username] = Account{Username: username, Password: password}
	fmt.Fprintf(w, "Account created successfully")
}

func login(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	account, exists := accounts[username]
	if !exists || account.Password != password {
		fmt.Fprintf(w, "Invalid username or password")
		return
	}

	fmt.Fprintf(w, "Logged in successfully")
}
