package common

var (
	MENUITEM_QUIT     = &MenuChoice{"B", "(B) Beenden"}
	MENUITEM_LOGIN    = &MenuChoice{"L", "(L) Login"}
	MENUITEM_REGISTER = &MenuChoice{"R", "(R) Registrieren"}

	MENUITEM_CHANGE_PASS    = &MenuChoice{"P", "(P) Passwort ändern"}
	MENUITEM_DELETE_ACCOUNT = &MenuChoice{"E", "(E) Eigenes Benutzerkonto löschen"}
	MENUITEM_LOGOUT         = &MenuChoice{"L", "(L) Logout"}
)

////////////////////////////////////////////////////////////////////////////////////////////////////

type Menu struct {
	Headertext  string
	MenuChoices []*MenuChoice
}

type MenuChoice struct {
	Shortcut string
	Text     string
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var MENU_MAIN = &Menu{
	"Startmenü",
	[]*MenuChoice{MENUITEM_LOGIN, MENUITEM_REGISTER, MENUITEM_QUIT},
	// fmt.Println("logout - logout")
	// fmt.Println("cpw - change current password") //TODO show and allow only if logged in
	// fmt.Println("del - delete account")          //TODO show and allow only if logged in
}

var MENU_LOGGED_IN = &Menu{
	"Hauptmenü",
	[]*MenuChoice{MENUITEM_CHANGE_PASS, MENUITEM_DELETE_ACCOUNT, MENUITEM_LOGOUT},
}
