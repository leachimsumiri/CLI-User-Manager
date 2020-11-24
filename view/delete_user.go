package view

import (
	"fmt"
)

func DeleteUser(username string) bool {
	ShowMessage("Möchtest Du den User '%s' wirklich löschen? (j/n)", username)
	var input string
	fmt.Scanln(&input)

	if input == "j" {
		return true
	}

	return false
}
