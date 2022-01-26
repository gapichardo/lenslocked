package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	// var m map[int]string
	mapa := map[int]string {
		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}
	data := struct {
		Name string
		Rol string
		Salary float32
		MiMapa map[int]string
	}{"John Smith","Manager",12344.56,mapa}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}