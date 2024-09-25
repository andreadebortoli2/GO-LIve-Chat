package helpers

import (
	"log"
	"net/http"

	"github.com/andreadebortoli2/GO-Live-Chat/internal/render"
)

func RenderErr(err error, w http.ResponseWriter, r *http.Request, page string, sm map[string]string) {
	log.Println(err)
	strErr := err.Error()
	render.RenderPage(w, r, page, render.TemplateData{
		StringMap: sm,
		Error:     strErr,
	})
}
