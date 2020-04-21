package database

import (
	"strconv"
	"time"
)

func convToDbKind(nachname, vorname, strAlter, strasseNr, plzOrt, schule, klasse, nachnameErz, vornameErz, email, telefon, strEssen string) dbKind {
	id := getAnzahlKinder()
	alter, _ := strconv.ParseInt(strAlter, 10, 64)
	essen, _ := strconv.ParseInt(strEssen, 10, 64)

	kind := dbKind{
		id,
		nachname,
		vorname,
		int(alter),
		strasseNr,
		plzOrt,
		schule,
		klasse,
		nachnameErz,
		vornameErz,
		email,
		telefon,
		int(essen - 1),
		0,
		false,
		time.Time{},
		false,
		time.Time{},
	}

	return kind
}
