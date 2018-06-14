package gsuiserver

import (
	"html/template"
	"net/http"
)

// StatusPageData is a structure of the dynamically loaded
// contents in the status page template.
type StatusPageData struct {
	PageTitle string
}

func statusPageResource(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	page := template.Must(template.ParseFiles("./templates/status.html"))
	data := StatusPageData{
		PageTitle: "Testing Templates",
	}
	page.Execute(w, data)
}
