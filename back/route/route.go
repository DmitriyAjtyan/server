package route

import (
	"fmt"
	"html/template"
	"net/http"
	config "server/back/config"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// RouterFunc is a main routing func of project
func RouterFunc() {

	r := mux.NewRouter()

	r.HandleFunc("/", homePageHandler)
	r.HandleFunc("/api/first", apiFirstHandler)

	http.Handle("/", r)
}

// HomePageHandler func is a handler for "/" route
func homePageHandler(w http.ResponseWriter, r *http.Request) {
	var secretKey = []byte(config.ConfigData.JWTSecret)
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, _ := token.SignedString(secretKey)
	t, err := template.ParseFiles("../front/index.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	if err := t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), 400)
	}
	w.Write([]byte(tokenString))
}

// ApiFirstHandler func is a handler for "/api/first" route
func apiFirstHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userJWTstring := r.Form["jwt"][0]
	token, err := jwt.Parse(userJWTstring, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.ConfigData.JWTSecret), nil
	})
	if err != nil {
		fmt.Print("JWT token parse error: ", err)
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		t, err := template.ParseFiles("../front/apiFirst.html")
		if err != nil {
			http.Error(w, err.Error(), 400)
		}
		if err := t.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), 400)
		}
	} else {
		t, err := template.ParseFiles("../front/apiSecond.html")
		if err != nil {
			http.Error(w, err.Error(), 400)
		}
		if err := t.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), 400)
		}
	}
}
