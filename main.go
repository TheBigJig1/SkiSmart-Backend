package main

import (
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
)

type User struct {
	Email       string
	Password    string
	First, Last string
	Zipcode     string
}

var (
	users = map[string]User{}
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/user/create", Create)
	mux.HandleFunc("/user/login", Login)
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Println("Listening...")
	server.ListenAndServe() // Run the http server
}

func Create(w http.ResponseWriter, r *http.Request) {
	slog.Info("recieved create request")
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u := User{}
	u.Email = r.FormValue("email")
	u.Password = r.FormValue("password")
	u.First = r.FormValue("first")
	users[u.Email] = u
}

func Login(w http.ResponseWriter, r *http.Request) {
	slog.Info("recieved create request")
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	email := r.FormValue("email")
	passwd := r.FormValue("password")

	user, ok := users[email]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if user.Password != passwd {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&user)
}
