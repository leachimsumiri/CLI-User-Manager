package view

import (
	"fmt"
	"strings"
)

//func get_input(menu map[string]int, dont_return_util_correct_pseudo_optional_param ...bool) int { // Evtl optionale "Abbrechmöglichkeit" übergeben können
func get_input(menu map[string]int) int {

	var input string

	for {
		fmt.Scanln(&input)
		for target_key, target_id := range menu {
			if strings.ToUpper(target_key) == strings.ToUpper(input) {
				return target_id
			}
		}
		// fmt.Printf("\r") // Vorherige Eingabe löschen (carriage return) - Wegen Zeilenwechsel aktuell nicht sinnvoll
	}
}
