package guiserver

import (
	"log"
	"net/http"

	"github.com/d3coy/ESRA_GroundStation/models"
	"github.com/jinzhu/gorm"
)

const (
	serverAddress = ":10080"
	dbPath        = "./flight.db"
)

// Env is an environment map structure passed between handle functions.
type Env struct {
	DB *gorm.DB
}

// GroundStationUI is an http webserver that serves dynamically
// generated templates to the end user.
func GroundStationUI() {
	db := models.InitDB(dbPath)
	env := &Env{DB: db}

	fs := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL, r.Proto)
		http.FileServer(http.Dir("static")).ServeHTTP(w, r)
	})

	http.HandleFunc("/", indexResource)
	http.HandleFunc("/status/", env.StatusHandler)
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server now running on", serverAddress)
	if err := http.ListenAndServe(serverAddress, nil); err != nil {
		panic(err)
	}
}

func indexResource(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL, r.Proto)
	http.Redirect(w, r, "/status/", http.StatusFound)
}
