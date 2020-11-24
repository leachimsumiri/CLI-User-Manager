package view

import (
	"fmt"
)

func GetPassword(minLength int) string {
	ShowMessage("Passwort (min. %d Zeichen): ", minLength)
	var password string
	fmt.Scanln(&password)

	// if checkPasswordLength(password) == ERR_PASSWORD_SHORT {
	// 	return EMPTY_STRING
	// }

	return password
}
