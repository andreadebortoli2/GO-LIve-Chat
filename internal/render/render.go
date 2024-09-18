package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func addData(dataToAdd map[string]string) map[string]string {
	data := map[string]string{}
	data["test"] = "passing data to page from render"
	for key, content := range dataToAdd {
		data[key] = content
	}

	return data
}

// RenderPage render the requested page
func RenderPage(w http.ResponseWriter, r *http.Request, pageName string, handlerData map[string]string) {

	page, err := pagesCache()
	if err != nil {
		log.Println(err)
		return
	}

	data := addData(handlerData)

	page[fmt.Sprintf("%s.page.html", pageName)].Execute(w, data)
}

// pagesCache create a map with all the pages parsed with the layouts and saved with their name as key
func pagesCache() (map[string]*template.Template, error) {

	parsePages := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		log.Println("can't find pages files")
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		parsePage, err := template.ParseFiles(page)
		if err != nil {
			log.Println("can't parse the page", err)
			return nil, err
		}
		parsePage, err = parsePage.ParseFiles("./templates/base.layout.html")
		if err != nil {
			log.Println("can't parse the layout", err)
			return nil, err
		}
		parsePages[name] = parsePage
	}

	return parsePages, nil
}
