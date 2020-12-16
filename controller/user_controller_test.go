package controller

import (
	"testing"

	"gitlab.com/fh-campus/sde22-asd-exercise/model"
)

//USERCONTROLLER TESTS, JE EIN GUTFALL UND EIN SCHLECHTFALL

const (
	TESTFIRSTNAME       = "TestFirstName"
	TESTLASTNAME        = "TestLastName"
	TESTUSERNAME        = "TestUserName"
	TESTPASSWORD        = "TestPassword"
	NONEXISTENTUSERNAME = "thisUserShouldNotExist"
	INCORRECTPASSWORD   = "incorrectPassword"
)

/*
- makes sure testuser doesn't exist
- makes sure there is no User logged in (technically not necessary yet due to current architecture)
- registers a TestUser
*/
func createPreconditions(t *testing.T, userController *UserController) {
	if userController.loggedInUser != nil || userController.Exists(TESTUSERNAME) != false {
		userController.loggedInUser = userController.GetByUsername(TESTUSERNAME)
		model.Inst.Unscoped().Delete(userController.loggedInUser)
		userController.loggedInUser = nil
	}

	exists := userController.Exists(TESTUSERNAME)
	if exists == true {
		t.Errorf("Test Error. User can not exist before creation")
	}

	userController.RegisterUser(TESTFIRSTNAME, TESTLASTNAME, TESTUSERNAME, TESTPASSWORD)
}

/*
- logs out
- deletes testuser
*/
func teardown(userController *UserController) {
	exists := userController.Exists(TESTUSERNAME)
	isLoggedIn := userController.loggedInUser != nil

	if isLoggedIn == true {
		userController.loggedInUser = nil
	}

	if exists == true {
		userController.loggedInUser = userController.GetByUsername(TESTUSERNAME)
		model.Inst.Unscoped().Delete(userController.loggedInUser) //hard delete
	}
}

func TestCheckPassword(t *testing.T) {
	var userController *UserController = &UserController{}

	createPreconditions(t, userController)

	user := userController.GetByUsername(TESTUSERNAME)

	correctPassword := userController.checkPassword(user, TESTPASSWORD)
	if correctPassword == false {
		t.Errorf("Test Error. Password check with correct password returns false")
	}

	incorrectPassword := userController.checkPassword(user, INCORRECTPASSWORD)
	if incorrectPassword == true {
		t.Errorf("Test Error. Password check with incorrect password returns true")
	}

	teardown(userController)
}

func TestGetByUsername(t *testing.T) {
	var userController *UserController = &UserController{}

	createPreconditions(t, userController)

	exists := userController.GetByUsername(TESTUSERNAME)
	if exists == nil {
		t.Errorf("Test Error. Testuser does not exist after creation")
	}

	shouldNotExist := userController.GetByUsername(NONEXISTENTUSERNAME)
	if shouldNotExist != nil {
		t.Errorf("Test Error. Non existent user returns exist")
	}

	teardown(userController)
}

func TestIsLoggedIn(t *testing.T) {
	var userController *UserController = &UserController{}

	createPreconditions(t, userController)

	userController.TryLoginUser(TESTUSERNAME, TESTPASSWORD)
	isLoggedIn := userController.loggedInUser != nil
	if isLoggedIn == false {
		t.Errorf("Test Error. Login with correct Credentials failed")
	}

	userController.loggedInUser = nil

	userController.TryLoginUser(TESTUSERNAME, "WRONG PASSWORD TEST")
	isNotLoggedIn := userController.loggedInUser == nil
	if isNotLoggedIn == false {
		t.Errorf("Test Error. Login with wrong Credentials went through")
	}

	teardown(userController)
}

func TestUserExists(t *testing.T) {
	var userController *UserController = &UserController{}

	createPreconditions(t, userController)

	exists := userController.Exists(TESTUSERNAME)
	if exists == false {
		t.Errorf("Test Error. User did not exist after creation")
	}

	shouldNotExist := userController.Exists(NONEXISTENTUSERNAME)
	if shouldNotExist == true {
		t.Errorf("Test Error. User %s should not exist.", NONEXISTENTUSERNAME)
	}

	teardown(userController)
}
