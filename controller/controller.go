package controller

import (
	"os"
	"time"

	"gitlab.com/fh-campus/sde22-asd-exercise/common"
	"gitlab.com/fh-campus/sde22-asd-exercise/model"
	"gitlab.com/fh-campus/sde22-asd-exercise/view"
)

const (
	//REQUEST_LOGOUT = iota
	TIMEOUT_LOGGED_IN = iota
	TIMEOUT_LOGGED_OUT
)

const LOGIN_TIMEOUT_IN_SECONDS = 60
const EXIT_CODE_TIMEOUT = 42

func Main() {

	var err error
	requestChannel := make(chan int)
	timeoutChannel := make(chan int)

	go timeoutHandler(timeoutChannel, requestChannel)

	view.ShowWelcome()
	for {

		// Aktuelles Menü anzeigen
		if userController.IsLoggedIn() {
			timeoutChannel <- TIMEOUT_LOGGED_IN // Aktivität / Timer rücksetzen
			err = MainMenu(common.MENU_LOGGED_IN, timeoutChannel)
		} else {
			err = MainMenu(common.MENU_MAIN, timeoutChannel)
		}

		// Fehler ausgeben
		if err != nil {
			view.ShowMessage("%v", err)
		}
	}
}

func quit() {
	view.ShowGoodbye()
	os.Exit(0)
}

func MainMenu(menu *common.Menu, timeoutCh chan<- int) (err error) {
	switch view.ShowMenu(menu) {
	case common.MENUITEM_LOGIN:
		username, password := view.GetLogin()
		err = userController.TryLoginUser(username, password)

	case common.MENUITEM_REGISTER:
		firstName, lastName, username, password := view.AddUser(model.MIN_PASSWORD_LENGTH)
		err = userController.RegisterUser(firstName, lastName, username, password)

	case common.MENUITEM_CHANGE_PASS:
		userController.ChangePassword()

	case common.MENUITEM_DELETE_ACCOUNT:
		userController.DeleteUser()

	case common.MENUITEM_LOGOUT:
		timeoutCh <- TIMEOUT_LOGGED_OUT // only send once (= when logging out)
		userController.Logout()

	case common.MENUITEM_QUIT:
		quit()
	}

	return
}

func timeoutHandler(inchan chan int, outchan chan int) {
	var logInState = TIMEOUT_LOGGED_OUT
	timer := time.NewTimer(time.Hour)
	timer.Stop()
	for {
		// Reset timer whenever activity (+ stop when logged out)
		select {
		case logInState = <-inchan:
			switch logInState {
			case TIMEOUT_LOGGED_OUT:
				timer.Stop()
			case TIMEOUT_LOGGED_IN:
				timer.Reset(LOGIN_TIMEOUT_IN_SECONDS * time.Second)
			}
		case <-timer.C:
			view.ShowMessage("Ausgeloggt wegen Inaktivität")
			os.Exit(EXIT_CODE_TIMEOUT) // wegen Consolenanwendung unaufwendige Möglichkeit (ganz beenden)
			//outchan <- REQUEST_LOGOUT // eigentlich von dieser Go-Routine nur rückmelden
		}
	}
}
