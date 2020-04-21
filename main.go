package main

import (
	"html/template"
	"log"
	"net/http"
	//"github.com/landupin/kinderbibeltag/anmelden"
	//"github.com/landupin/kinderbibeltag/db"
)

/*Data for html templates*/
type Data struct {
	Kibitagsets      Settings
	Error            bool
	SubKind          formKind
	ValidationErrors formKind
	AngemeldetKind   formKind
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("public"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		anmelden(w, r)
		return
	}

	err := tpl.ExecuteTemplate(w, "index.html", Data{Kibitagsets: kibitagset})
	HandleError(w, err)
	return
}

/*HandleError ...*/
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}
}
