package route

import (
	"html/template"
	"net/http"

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
	t, err := template.ParseFiles("../front/index.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	if err := t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), 400)
	}
}

// ApiFirstHandler func is a handler for "/api/first" route
func apiFirstHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../front/apiFirst.html")
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
	if err := t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), 400)
	}
}
