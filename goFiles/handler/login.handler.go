package goFiles

func loginUser(w http.ResponseWriter, r *http.Request) {
	type LoginResponse struct {
		currentlyLoggedInUserName string
		err string
	}

	type UserLoggIn struct {
		currentlyLoggedInUserName string
		errorWrongUsername        error
	}

	t, err := template.ParseFiles("htmlTemplates/login.html")
	if err != nil {
		fmt.Println(err)
	}

	// If input = logout, the user wants to log out
	// logout information comes through the "name" fild in login

	currentlyLoggedInUser := r.FormValue("usernameLogin")
	// not a nice solution using a string called logout1 instead of a bool, but works currently
	usernameLogout := r.FormValue("logout1")

	currentlyUser := UserLoggIn{
		currentlyLoggedInUserName: currentlyLoggedInUser,
		errorWrongUsername:        invalidUsername(),
	}

	if usernameLogout == "logout" {
		currentlyUser.currentlyLoggedInUserName = ""
		fmt.Println(currentlyUser.currentlyLoggedInUserName)
	}

	// try to log in user
	if checkIfUserExists(currentlyLoggedInUser) {
		Login(currentlyLoggedInUser)
		// testing because of logout in the html partcurrentlyLoggedInUserName login
		t.Execute(w, currentlyUser)

		// response := LoginResponse{currentlyLoggedInUserName: currentlyLoggedInUser, err: ""}
		// t.Execute(w, response)

		// Login without an valid username
	} else {
		// response := LoginResponse{currentlyLoggedInUserName: "", err: "Deine Anmeldedaten sind ungültig"}
		// currentlyUser.errorWrongUsername = t.Execute(w, response)
		// currentlyLoggedInUser = ""

		currentlyUser.errorWrongUsername =
			t.Execute(w, currentlyUser.errorWrongUsername)
		currentlyLoggedInUser = ""
	}
}
