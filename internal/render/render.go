package render

import (
	"fmt"
	"html/template"
	"log"
)

func RenderPage(pageName string) (*template.Template, error) {
	parsePage, err := template.ParseFiles(fmt.Sprintf("./templates/%s.page.html", pageName))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	parsePage, err = parsePage.ParseFiles("./templates/base.layout.html")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return parsePage, nil
}
