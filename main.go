package main

import (
	_ "fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {

	router := http.NewServeMux()

	router.HandleFunc("/{$}", handleRoot)
	router.HandleFunc("/auth/login", handleLogin)
	router.HandleFunc("/auth/signup", handleSignup)

	log.Println("Listning")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}

func handleRoot(w http.ResponseWriter, _ *http.Request) {
	index := template.Must(template.ParseFiles("templates/index.html"))
	index.Execute(w, nil)
}

func handleLogin(w http.ResponseWriter, _ *http.Request) {
	login := template.Must(template.ParseFiles("templates/login.html"))
	login.Execute(w, nil)
}

func handleSignup(w http.ResponseWriter, _ *http.Request) {
	signup := template.Must(template.ParseFiles("templates/signup.html"))
	signup.Execute(w, nil)
}
