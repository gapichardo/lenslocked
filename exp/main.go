package main

import (
	"fmt"

	_ "github.com/lib/pq"
	"lenslocked.com/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "picge02"
	dbname   = "lenslockedDB"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname)

	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	us.DestructiveReset()

	// Create a user
	user := models.User{
		Name:     "Michael",
		LastName: "Scott",
		Email:    "michael@dundermifflin.com",
	}

	user2 := models.User{
		Name:           "Gerardo",
		LastName:       "Pichardo",
		SecondLastName: "Forcade",
		Email:          "gerardo.pichardo@dundermifflin.com",
	}

	if err := us.Create(&user); err != nil {
		panic(err)
	}
	if err := us.Create(&user2); err != nil {
		panic(err)
	}

	foundUser, err := us.ByEmail("michael@dundermifflin.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(foundUser)
	// Update a user
	user.Name = "Updated Name"
	if err := us.Update(&user); err != nil {
		panic(err)
	}
	foundUser2, err2 := us.ByEmail("michael@dundermifflin.com")
	if err2 != nil {
		panic(err2)
	}
	// Because of an update, the name should now be "Updated Name"
	fmt.Println(foundUser2)

	foundUser3, err3 := us.ByLastName("Pichardo")
	if err3 != nil {
		panic(err3)
	}
	fmt.Println(foundUser3)
}
