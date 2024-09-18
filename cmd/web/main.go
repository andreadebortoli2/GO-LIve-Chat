package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/", ShowPage)

	fmt.Println("serving on port 8080")
	_ = http.ListenAndServe(":8080", nil)

}

func ShowPage(w http.ResponseWriter, r *http.Request) {

	parsePage, err := template.ParseFiles("./templates/index.layout.html")
	if err != nil {
		log.Println(err)
		return
	}
	parsePage, err = parsePage.ParseFiles("./templates/home.page.html")
	if err != nil {
		log.Println(err)
		return
	}

	parsePage.Execute(w, nil)
}
