package view

import (
	"fh-asd-2/common"
	"fmt"
	"strings"
)

// DOEV: Könnte man noch aufbrechen in menu_output(*Menu) + menu_input(*Menu) um sich Kommentare zu ersparen (aber evtl zu "optimiert")
func ShowMenu(menu *common.Menu) *common.MenuChoice {

	// Output menu text
	fmt.Println("")
	fmt.Println(strings.Repeat("-", len(menu.Headertext)))
	fmt.Println(menu.Headertext)
	fmt.Println(strings.Repeat("-", len(menu.Headertext)))
	for _, choice := range menu.MenuChoices {
		fmt.Println(choice.Text)
	}

	// Input user response
	var input string
	for {
		fmt.Scanln(&input)
		for _, choice := range menu.MenuChoices {
			if strings.ToUpper(choice.Shortcut) == strings.ToUpper(input) {
				return choice
			}
		}
	}

}
