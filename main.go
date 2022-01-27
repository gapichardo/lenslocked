package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"lenslocked.com/controllers"
)

func main() {
	usersC := controllers.NewUsers()
	staticC := controllers.NewStatic()
	galleryC := controllers.NewGalleries()

	var h http.Handler = http.Handler(staticC.NotFound)

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/faq", staticC.Faq).Methods("GET")

	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.HandleFunc("/galleries/new", galleryC.New).Methods("Get")
	r.HandleFunc("/galleries/new", galleryC.Create ).Methods("Post")

	r.NotFoundHandler = h
	http.ListenAndServe(":3000", r)
}
