package controller

import (
	"errors"
	"fh-asd-2/model"
	"fh-asd-2/view"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ERR_PASSWORD_SHORT         = errors.New("Gewähltes Passwort zu kurz")
	ERR_USERNAME_EXISTS        = errors.New("Username existiert bereits")
	ERR_USERNAMEPASS_INCORRECT = errors.New("username oder password nicht korrekt")
	EMPTY_STRING               = ""
)

const MAX_LOGIN_TRIES = 3

////////////////////////////////////////////////////////////////////////////////////////////////////

var userController *UserController = &UserController{}

type UserController struct {
	loggedInUser *model.User // nil => nicht eingeloggt
	loginTries   int
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func (uc *UserController) IsLoggedIn() bool {
	return uc.loggedInUser != nil
}

func (*UserController) GetByUsername(username string) *model.User {
	user := &model.User{}
	err := model.Inst.Where("Username = ?", username).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}
	if err != nil {
		log.Fatalf("Fataler Fehler in GetByUsername: %v", err)
	}
	return user
}

func (uc *UserController) Exists(username string) bool {
	return uc.GetByUsername(username) != nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func (uc *UserController) hashPassword(plaintextPassword string) (hashedPassword string) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), bcrypt.MinCost)
	if err != nil { // "darf eigentlich nicht passieren"
		log.Fatalf("Fataler Fehler beim Erstellen des Passworthashes: %v", err)
	}
	return string(hashedBytes)
}

func (uc *UserController) checkPassword(user *model.User, plaintextPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(plaintextPassword))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false
	}
	if err != nil {
		log.Fatalf("Fataler Fehler beim Vergleichen des Passworthashes: %v", err)
	}
	return true
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func (this *UserController) TryLoginUser(usernameInput string, passwordInput string) error {
	user := userController.GetByUsername(usernameInput)
	if user == nil || !this.checkPassword(user, passwordInput) { //user.Password != this.Hash(password) {
		this.loginTries++
		if this.loginTries >= MAX_LOGIN_TRIES {
			// Zu viele Versuche => Software beenden
			log.Fatalf("Zu viele inkorrekte Loginversuche!") // TODO: Anforderung nicht völlig klar => sperren für 10 Minuten?
		}
		return ERR_USERNAMEPASS_INCORRECT
	}

	// Erfolgreich angemeldet
	this.loggedInUser = user
	view.ShowMessage("Benutzer \"%s\" erfolgreich eingeloggt", user.Username)
	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func (this *UserController) RegisterUser(firstName string, lastName string, username string, password string) error {

	if len(password) < model.MIN_PASSWORD_LENGTH {
		return ERR_PASSWORD_SHORT
	}
	if userController.Exists(username) {
		return ERR_USERNAME_EXISTS
	}

	// Benutzer anlegen
	user := &model.User{FirstName: firstName, LastName: lastName, Username: username, Password: this.hashPassword(password)}
	res := model.Inst.Create(user)
	if res.Error != nil {
		// Programm beenden/killen (bei jedwegem unerwarteten Fehler)
		log.Fatalf("Datenbankfehler beim Benutzeranlegen: %+v", res.Error)
	}

	// Kein Fehler
	view.ShowMessage("Benutzer \"%s\" erfolgreich erstellt", username)
	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func (this *UserController) Logout() {
	this.loggedInUser = nil
	view.ShowMessage("Benutzer ausgeloggt.")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func (this *UserController) DeleteUser() {
	response := view.DeleteUser(this.loggedInUser.Username)

	if response {
		model.Inst.Unscoped().Delete(&this.loggedInUser) // unscoped => permanently, no soft delete!
		view.ShowMessage("User '%s' gelöscht.", this.loggedInUser.Username)
		this.loggedInUser = nil
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func (this *UserController) ChangePassword() {
	password := EMPTY_STRING

	for password == EMPTY_STRING {
		password = view.GetPassword(model.MIN_PASSWORD_LENGTH)
	}

	if len(password) < model.MIN_PASSWORD_LENGTH {
		view.ShowMessage("Passwort zu kurz")
		return
	}

	this.loggedInUser.Password = this.hashPassword(password)
	model.Inst.Save(this.loggedInUser)
	view.ShowMessage("Passwort erfolgreich geändert")
}
