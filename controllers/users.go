package controllers

import (
	"fmt"
	"net/http"

	"lenslocked.com/models"
	"lenslocked.com/views"
)

type Users struct {
	NewView *views.View
	us      *models.UserService
}

func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
		us:      us,
	}
}

type SignupForm struct {
	Name           string `schema:"name"`
	LastName       string `schema:"lastname"`
	SecondLastName string `schema:"secondlastname"`
	Email          string `schema:"email"`
	Password       string `schema:"password"`
}

// This function execute when we call th new.gohtml template with the GET method
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// Create is used to process the signup form when a user tries to create a new user account.
//
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user := models.User{
		Name:           form.Name,
		LastName:       form.LastName,
		SecondLastName: form.SecondLastName,
		Email:          form.Email,
	}
	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "User is:", user)
}
