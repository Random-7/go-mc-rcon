package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}

}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	http.ListenAndServe(":8080", nil)

}

// rcon.SendCommand("list")
// rcon.SendCommand("tps")
// rcon.SendCommand("seed")

// connectionInfo := new(database.DbConnection)
// connectionInfo.Dsn = "go-mc-rcon:Fingas@1437@tcp(10.0.10.5:3306)/go-mc-rcon?charset=utf8mb4&parseTime=True&loc=Local"
// database.Connect(*connectionInfo)
