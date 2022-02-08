package controllers

import (
	"fmt"
	"net/http"

	"lenslocked.com/views"
)

func NewGalleries() *Galleries {
	return &Galleries{
		GalleryView: views.NewView("bootstrap", "gallery/newGallery"),
	}
}

type Galleries struct {
	GalleryView *views.View
}

type GalleryForm struct {
	Gallery   string `schema:"gallery"`
	Tipo      string `schema:"tipo"`
	Categoria string `schema:"categoria"`
}

// This function execute when we call th new.gohtml template with the GET method
// GET /signup
func (g *Galleries) New(w http.ResponseWriter, r *http.Request) {
	if err := g.GalleryView.Render(w, nil); err != nil {
		panic(err)
	}
}

// Create is used to process the signup form when a user
// tries to create a new user account.
//
// POST /signup
func (g *Galleries) Create(w http.ResponseWriter, r *http.Request) {
	var form GalleryForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, "Name is", form.Gallery)
	fmt.Fprintln(w, "Tipo is", form.Tipo)
	fmt.Fprintln(w, "Tipo is", form.Categoria)
}
