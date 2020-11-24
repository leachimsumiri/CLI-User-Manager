package view

import (
	"fmt"
)

func GetLogin() (username string, password string) {
	fmt.Print("Username: ")
	fmt.Scanln(&username)

	fmt.Print("Passwort: ")
	fmt.Scanln(&password)

	return // gibt die in der Definition angegeben Parameter zurück (username, password)
}
