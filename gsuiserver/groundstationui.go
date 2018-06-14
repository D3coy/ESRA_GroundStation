package gsuiserver

import (
	"log"
	"net/http"
)

const (
	address = ":10080"
)

// GroundStationUI is an http webserver that serves dynamically
// generated templates to the end user.
func GroundStationUI() {
	fs := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		http.FileServer(http.Dir("static")).ServeHTTP(w, r)
	})

	http.HandleFunc("/", indexPage)
	http.HandleFunc("/status/", statusPageResource)
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server now running on", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		panic(err)
	}
}

func logRequest(r *http.Request) {
	log.Println(r.Method, r.URL, r.Proto)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	http.Redirect(w, r, "/status/", 302)
}
