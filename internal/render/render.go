package render

import (
	"html/template"
	"log"
)

func RenderPage() (*template.Template, error) {
	parsePage, err := template.ParseFiles("./templates/index.layout.html")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	parsePage, err = parsePage.ParseFiles("./templates/home.page.html")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return parsePage, nil
}
