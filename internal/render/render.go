package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/andreadebortoli2/GO-Experiment-and-Learn/internal/config"
	"github.com/justinas/nosurf"
)

var appConfig *config.AppConfig

func NewRenderer(a *config.AppConfig) {
	appConfig = a
}

// addData add data to template
func addData(dataToAdd map[string]string, r *http.Request) map[string]string {
	data := map[string]string{}
	data["CSRFToken"] = nosurf.Token(r)
	data["auth"] = appConfig.Session.GetString(r.Context(), "auth")

	return data
}

// RenderPage render the requested page
func RenderPage(w http.ResponseWriter, r *http.Request, pageName string, handlerData map[string]string) error {

	pages, _ := pagesCache()

	requestedPage, ok := pages[fmt.Sprintf("%s.page.html", pageName)]
	if !ok {
		return errors.New("page not found in cache")
	}

	data := addData(handlerData, r)

	buf := new(bytes.Buffer)

	err := requestedPage.Execute(buf, data)
	if err != nil {
		return err
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		return err
	}

	return nil
}

// pagesCache create a map with all the pages parsed with the layouts and saved with their name as key
func pagesCache() (map[string]*template.Template, error) {

	pagesCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		log.Println("can't find pages files")
		return pagesCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		parsePage, err := template.ParseFiles(page)
		if err != nil {
			log.Println("can't parse the page", err)
			return pagesCache, err
		}
		parsePage, err = parsePage.ParseFiles("./templates/base.layout.html")
		if err != nil {
			log.Println("can't parse the layout", err)
			return pagesCache, err
		}
		pagesCache[name] = parsePage
	}

	return pagesCache, nil
}
