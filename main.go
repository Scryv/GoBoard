package main

import (
	_ "fmt"
	"github.cm/golang-jwt/jwt/v5"
	"html/template"
	"log"
	"net/http"
	"time"
)

var secretKey = []byte("SuperSecretKeyFr") //here as placeholder will be more secure later

func main() {

	router := http.NewServeMux()

	router.HandleFunc("/{$}", handleRoot)
	router.HandleFunc("/auth/login", handleLogin)
	router.HandleFunc("/auth/signup", handleSignup)
	router.HandleFunc("/login", login)
	router.HandleFunc("/signup", signup)
	router.HandleFunc("/error", erreur)

	log.Println("Listning")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}

func handleRoot(w http.ResponseWriter, _ *http.Request) {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		log.Println("No Cookie Found")
		return
	}
	err = verifyToken(cookie.Value)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}
	fmt.Fprintf(w, "Excuse me sir there has been something you have confused me for")
	fmt.Fprintf(w, "Cookie is Found: %s = %s", cookie.Name, cookie.Value)
	index := template.Must(template.ParseFiles("templates/index.html"))
	index.Execute(w, nil)
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
	log.Println("Trying to log in as: ", username)
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
	fmt.Fprintln(w, "Cookie has been set!")
}

func handleSignup(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	log.Println("These where used: ", username, password)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
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
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
