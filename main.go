// main.go
package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("your-secret-key"))

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/login", LoginHandler).Methods("GET", "POST")
	r.HandleFunc("/signup", SignupHandler).Methods("GET", "POST")
	r.HandleFunc("/logout", LogoutHandler).Methods("GET")
	r.HandleFunc("/todos", TodoListHandler).Methods("GET")
	r.HandleFunc("/todos/new", NewTodoHandler).Methods("GET", "POST")
	r.HandleFunc("/todos/{id:[0-9]+}/edit", EditTodoHandler).Methods("GET", "POST")
	r.HandleFunc("/todos/{id:[0-9]+}/delete", DeleteTodoHandler).Methods("POST")

	http.Handle("/", r)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
