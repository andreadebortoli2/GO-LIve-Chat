package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/andreadebortoli2/GO-Live-Chat/internal/config"
	"github.com/andreadebortoli2/GO-Live-Chat/internal/models"
	"github.com/justinas/nosurf"
)

var appConfig *config.AppConfig

func NewRenderer(a *config.AppConfig) {
	appConfig = a
}

type TemplateData struct {
	StringMap  map[string]string
	Error      string
	Data       map[string]interface{}
	CSRFToken  string
	ActiveUser ActiveUser
}

type ActiveUser struct {
	ID          uint
	UserName    string
	Email       string
	AccessLevel int
}

func addDataToTemplate(td TemplateData, r *http.Request) TemplateData {
	td.CSRFToken = nosurf.Token(r)

	ses, err := appConfig.Session.Get(r, "active_user")
	if err != nil {
		log.Println("cannot get session (render add data)", err)
		return td
	}

	if !ses.IsNew {
		u := ses.Values["user"].(models.User)
		td.ActiveUser.ID = u.ID
		td.ActiveUser.UserName = u.UserName
		td.ActiveUser.Email = u.Email
		accLvl, err := strconv.Atoi(u.AccessLevel)
		if err != nil {
			log.Println("cannot convert access level")
			td.ActiveUser.AccessLevel = 0
		}
		td.ActiveUser.AccessLevel = accLvl
	}

	return td
}

// RenderPage render the requested page
func RenderPage(w http.ResponseWriter, r *http.Request, pageName string, handlerData TemplateData) error {

	pages, _ := pagesCache()

	requestedPage, ok := pages[fmt.Sprintf("%s.page.html", pageName)]
	if !ok {
		return errors.New("page not found in cache")
	}

	data := addDataToTemplate(handlerData, r)

	buf := new(bytes.Buffer)

	err := requestedPage.Execute(buf, data)
	if err != nil {
		log.Println("page execute err", err)
		return err
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println("write err:", err)
		return err
	}

	return nil
}

// pagesCache create a map with all the pages parsed with the layouts and saved with their name as key
func pagesCache() (map[string]*template.Template, error) {

	pagesCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		log.Println("can't find pages files", err)
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
