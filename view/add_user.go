package view

import (
	"fmt"
)

func AddUser(minPassLen int) (firstName string, lastName string, username string, password string) {

	fmt.Println("")
	fmt.Println("Neuen Benutzer anlegen")
	fmt.Println("----------------------")

	//evtl. immer ein Abbruchkriterium anbieten
	// Hier sehr simples Konsolenprogramm -> kann man abbrechen
	// Wenn wir es auf Webservice erweitern, kein Problem mehr!

	fmt.Print("Vorname: ") //blank for empty? / Optional?
	fmt.Scanln(&firstName)

	fmt.Print("Nachname: ") //blank for empty? / Optional?
	fmt.Scanln(&lastName)

	fmt.Print("Username: ")
	fmt.Scanln(&username)

	fmt.Printf("Passwort (min. %d Zeichen): ", minPassLen) // keine spezifizierte Anforderung, aber auch gut denk ich auch
	fmt.Scanln(&password)

	return
}
