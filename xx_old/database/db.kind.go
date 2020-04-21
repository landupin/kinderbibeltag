package database

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type dbKind struct {
	ID                  int
	Nachname            string
	Vorname             string
	Alter               int
	StrasseNr           string
	PlzOrt              string
	Schule              string
	Klasse              string
	NachnameErz         string
	VornameErz          string
	Email               string
	Telefon             string
	Essen               int
	Reservierungsnummer int
	Bezahlt             bool
	Bezahltdatum        time.Time
	Bestaetigt          bool
	Bestaetigtdatum     time.Time
}

/*Employee ... */
type employee struct {
	Name     string
	Role     string
	HireDate time.Time
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "i am online!!elf!")

	ctx := appengine.NewContext(r)

	e1 := employee{
		Name:     "Joe Citizen",
		Role:     "Manager",
		HireDate: time.Now(),
	}
	fmt.Fprintln(w, e1)

	key, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "employee", nil), &e1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var e2 employee
	if err = datastore.Get(ctx, key, &e2); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Stored and retrieved the Employee named %q", e2.Name)

}

func getAnzahlKinder() int {
	return rand.Int()
}
