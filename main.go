package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

var tkk []byte
var ygap string

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	tkk := []byte(os.Getenv("TKN_K"))
	ygap := os.Getenv("YGAP_K")

	log.Println(ygap)
	log.Println(tkk)
	router := http.NewServeMux()

	router.HandleFunc("/{$}", handleRoot)
	router.HandleFunc("/auth/login", handleLogin)
	router.HandleFunc("/auth/signup", handleSignup)
	router.HandleFunc("/login", login)
	router.HandleFunc("/signup", signup)
	router.HandleFunc("/error", erreur)

	log.Println("Listning")
	err = http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		log.Println("No Cookie Found")
		index := template.Must(template.ParseFiles("templates/index.html"))
		index.Execute(w, nil)
		return
	}
	err = verifyToken(cookie.Value)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}
	log.Println("Sir There has been a cookie detected")
}

func login(w http.ResponseWriter, _ *http.Request) {
	login := template.Must(template.ParseFiles("templates/login.html"))
	login.Execute(w, nil)
}

func signup(w http.ResponseWriter, _ *http.Request) {
	signup := template.Must(template.ParseFiles("templates/signup.html"))
	signup.Execute(w, nil)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	log.Println("Trying to log in as: ", username, password)
	log.Println(password)

	tkn, _ := createToken(username)
	cookie := http.Cookie{
		Name:     "access_token",
		Value:    tkn,
		MaxAge:   20,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}

	http.SetCookie(w, &cookie)
	fmt.Fprintln(w, "Cooker set")
}

func handleSignup(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	log.Println("These where used to make account: ", username, password)
	if len(password) > 0 && len(username) > 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(tkk)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func erreur(w http.ResponseWriter, _ *http.Request) {
	log.Println("TEST")
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return tkk, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
