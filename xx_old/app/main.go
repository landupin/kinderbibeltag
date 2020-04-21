// Sample datastore-quickstart fetches an entity from Google Cloud Datastore.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"google.golang.org/appengine"

	db "github.com/landupin/kinderbibeltag/database"
)

/*Data ...*/
type Data struct {
	Kibitagsets      db.Settings
	ValidationErrors formKind
	Error            bool
	SubKind          formKind
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("public"))))

	//http.HandleFunc("/put", put)
	http.HandleFunc("/get", get)

	appengine.Main()
	//log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		anmelden(w, r)
		return
	}

	settings, err := db.GetSettings(appengine.NewContext(r))
	if err != nil {
		HandleError(w, err)
		return
	}

	err = tpl.ExecuteTemplate(w, "index.html", Data{Kibitagsets: settings})
	HandleError(w, err)
	return
}

func get(w http.ResponseWriter, r *http.Request) {
	settings, err := db.GetSettings(appengine.NewContext(r))
	if err != nil {
		HandleError(w, err)
		return
	}

	fmt.Fprintln(w, time.Now())
	fmt.Fprintln(w, settings)

}

/*HandleError ...*/
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}
}

/*
INSERT THE SETTINGS TO DATASTORE
func insert(w http.ResponseWriter, r *http.Request) {
	loc, _ := time.LoadLocation("Europe/Berlin")
	kibitagset := db.Settings{
		Veranstaltungsname:  "Kinderbibeltag 2017",
		VeranstaltNameLang:  "Ökumenischer Kinderbibeltag Alfter",
		Seitenueberschrift:  "Anmeldung zum Ökumenischen Kinderbibeltag Alfter 2017",
		Veranstaltungsdatum: time.Date(2017, time.September, 30, 6, 0, 0, 0, loc),
		Anmeldestart:        time.Date(2017, time.September, 6, 6, 0, 0, 0, loc),
		AnmeldestartText:    "6. September 2017",
		Anmeldeschluss:      time.Date(2017, time.September, 22, 6, 0, 0, 0, loc),
		AnmeldeschlussText:  "22. September 2017",
		AnzahlKinderMax:     50,
		Mitbringen: []db.Mitbringsel{
			{1, "Kuchen", 8},
			{2, "Rohkost gewaschen und geschnitten (fuer 5 Kinder)", 7},
			{3, "Obst gewaschen und geschnitten (fuer 5 Kinder)", 6},
			{4, "2 Liter haltbare Milch", 4},
			{5, "2 Liter Apfelsaft", 4},
			{6, "5 Broetchen", 5},
			{7, "1 Tube Senf", 1},
			{8, "1 Tube Ketchup", 3},
			{9, "1 Paket Trinkkakaopulver", 1},
			{10, "Mitarbeit 11:30-12:30 Uhr bei der Essensvorbereitung im Kath. Pfarrheim", 2},
			{11, "Mitarbeit 12:30-13:30 Uhr bei Essensausgabe und Aufraeumen im Kath. Pfarrheim", 3},
			{12, "Mitarbeit 15:15-16:15 Uhr bei der Vorbereitung zum Kaffee im Evang. Gem.zentrum", 3},
			{13, "Mitarbeit 17:15-18:00 Uhr beim Aufraeumen im Evang. Gem.zentrum", 3},
		},
		Teilnahmebeitrag: 6.0,
		Bankverbindung:   "IBAN: DE94381602206501382424, BIC: GENODED1HBO bei der VR-Bank Bonn, Kontoinhaber: Dr. Eike Kohler",
	}

	if _, err := datastore.Put(appengine.NewContext(r), datastore.NewIncompleteKey(appengine.NewContext(r), "settings", nil), &kibitagset); err != nil {
		HandleError(w, err)
		return
	}
	fmt.Fprintln(w, "insert succeded")
}
*/
