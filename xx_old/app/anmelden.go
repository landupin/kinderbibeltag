package main

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	db "github.com/landupin/kinderbibeltag/database"
	"google.golang.org/appengine"
)

type formKind struct {
	Nachname      string
	Vorname       string
	Alter         string
	StrasseNr     string
	PlzOrt        string
	Schule        string
	Klasse        string
	Einverstanden bool
	NachnameErz   string
	VornameErz    string
	Email         string
	Telefon       string
	Essen         string
	DummeFrage    string
}

func anmelden(w http.ResponseWriter, r *http.Request) {
	//get POST data
	child := formKind{
		r.PostFormValue("nachname"),
		r.PostFormValue("vorname"),
		r.PostFormValue("alter"),
		r.PostFormValue("strasse_nr"),
		r.PostFormValue("plz_ort"),
		r.PostFormValue("schule"),
		r.PostFormValue("klasse"),
		r.PostFormValue("einverstanden") == "on",
		r.PostFormValue("nachname_erz"),
		r.PostFormValue("vorname_erz"),
		r.PostFormValue("email"),
		r.PostFormValue("telefon"),
		r.PostFormValue("essen"),
		r.PostFormValue("dumme_frage"),
	}

	fmt.Fprintln(w, "submitted: \n", child)

	//form validation
	berr, validationErrors, err := formValidations(child, r)
	if err != nil {
		HandleError(w, err)
	}
	//error handling of form validation
	if berr {
		settings, err := db.GetSettings(appengine.NewContext(r))
		if err != nil {
			HandleError(w, err)
		}

		d := Data{
			Kibitagsets:      settings,
			ValidationErrors: validationErrors,
			Error:            berr,
			SubKind:          child,
		}

		err = tpl.ExecuteTemplate(w, "index.html", d)
		HandleError(w, err)
		return
	}

	/*
		//db entry
		id := newDbKind(child)
		log.Print(child)

		//show page anmeldebestaetigung
		d := Data{
			Kibitagsets:    kibitagSets,
			AngemeldetKind: dbkinder[id],
		}

		tpl.ExecuteTemplate(w, "anmeldebestaetigung.html", d)
	*/
}

/* FORM VALIDATION METHOD */

func formValidations(child formKind, r *http.Request) (bool, formKind, error) {
	err := false
	validationErrors := formKind{}

	if child.Nachname == "" {
		validationErrors.Nachname = "Feld Nachname darf nicht leer sein"
		err = true
	}

	if child.Vorname == "" {
		validationErrors.Vorname = "Feld Vorname darf nicht leer sein"
		err = true
	}

	alter, _ := strconv.ParseInt(child.Alter, 10, 64)

	if child.Alter == "" {
		validationErrors.Alter = "Feld Alter darf nicht leer sein"
		err = true
	} else if reflect.TypeOf(int(alter)) != reflect.TypeOf(1) {
		validationErrors.Alter = "Ungültiger Wert für Alter"
		err = true
	}

	if child.StrasseNr == "" {
		validationErrors.StrasseNr = "Feld Strasse darf nicht leer sein"
		err = true
	}

	if child.PlzOrt == "" {
		validationErrors.PlzOrt = "Feld PLZ darf nicht leer sein"
		err = true
	}

	if child.Schule == "" {
		validationErrors.Schule = "Feld Schule darf nicht leer sein"
		err = true
	}

	if child.Klasse == "" {
		validationErrors.Klasse = "Feld Klasse darf nicht leer sein"
		err = true
	}

	validationErrors.Einverstanden = !child.Einverstanden
	if validationErrors.Einverstanden {
		err = true
	}

	if child.NachnameErz == "" {
		validationErrors.NachnameErz = "Feld Erz. Nachname darf nicht leer sein"
		err = true
	}

	if child.VornameErz == "" {
		validationErrors.VornameErz = "Feld Erz. Vorname darf nicht leer sein"
		err = true
	}

	if child.Email == "" {
		validationErrors.Email = "Feld E-Mail darf nicht leer sein"
		err = true
	} else if !strings.Contains(child.Email, "@") || !strings.Contains(child.Email, ".") {
		validationErrors.Email = "Ungültige E-Mail-Addresse"
		err = true
	}

	if child.Telefon == "" {
		validationErrors.Telefon = "Feld Telefonnr. darf nicht leer sein"
		err = true
	}

	int64indess, x := strconv.ParseInt(child.Essen, 10, 64)
	if x != nil {
		validationErrors.Essen = "Ungültige Eingabe"
		err = true
	}
	indess := int(int64indess) - 1
	if child.Essen == "0" || child.Essen == "" {
		validationErrors.Essen = "Feld Essen darf nicht leer sein"
		err = true
	} else {
		settings, x := db.GetSettings(appengine.NewContext(r))
		if x != nil {
			return err, validationErrors, errors.New("GetSettings failed")
		}
		if settings.Mitbringen[indess].Anzahl <= 0 {
			validationErrors.Essen = "Nicht mehr verfügbar"
			err = true
		}
	}

	if strings.Compare(child.DummeFrage, child.VornameErz) != 0 {
		validationErrors.DummeFrage = "Felder stimmen nicht überein"
		err = true
	}

	return err, validationErrors, nil
}
