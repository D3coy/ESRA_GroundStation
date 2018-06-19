package guiserver

import (
	"html/template"
	"log"
	"net/http"

	"github.com/d3coy/ESRA_GroundStation/models"
)

const (
	latLow   = 3221.0
	latHigh  = 3238.0
	lonLow   = -10689.0
	lonHigh  = -10663.0
	altitude = 20000
)

// StatusPageData is a structure of the dynamically loaded contents in the status page template.
type StatusPageData struct {
	PageTitle string
	Points    *[]models.GPSDataSet
	Axis      *graphAxis
}

type graphAxis struct {
	MinLat   float64
	MaxLat   float64
	MinLon   float64
	MaxLon   float64
	Altitude int
	Lat      int
	Lon      int
}

// StatusHandler contains the handle func for the status page.
func (env *Env) StatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL, r.Proto)
	p := template.Must(template.ParseFiles("./templates/status.template"))
	ds := models.GetGPSData(env.DB)
	ga := &graphAxis{
		MinLat:   latLow,
		MaxLat:   latHigh,
		MinLon:   lonLow,
		MaxLon:   lonHigh,
		Altitude: altitude,
		Lat:      int(latHigh - latLow),
		Lon:      int(lonHigh - lonLow),
	}

	d := StatusPageData{
		PageTitle: "Testing Templates",
		Points:    ds,
		Axis:      ga,
	}
	p.Execute(w, d)
}
