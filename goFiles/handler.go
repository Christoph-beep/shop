package goFiles

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"text/template"
)

//// Important globally ////

// for any mistakes
func error404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	t, _ := template.ParseFiles("htmlTemplates/404.html", "htmlTemplates/header.html", "htmlTemplates/footer.html")
	t.Execute(w, nil)

}

func (i Inventar) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

//// end ////

// parseGlob anschauen

func viewhandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Path[len("/view/"):]
	// fmt.Fprintf(w, "Hi there, I love %s!", username)
	inventar, err := loadInventory(username)
	if err != nil {
		http.Redirect(w, r, "/404/", http.StatusFound)
	}
	t, _ := template.ParseFiles("htmlTemplates/view.html")
	t.Execute(w, inventar)
}

// post request NICHT über die URl
// get über URL

func saveHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("userNameRegister")
	cobbleStoneString := r.FormValue("cobblestone")
	cobblestone, _ := strconv.Atoi(cobbleStoneString)
	p := Inventar{Username: username, Cobblestone: cobblestone}
	p.save()
	http.Redirect(w, r, "/view/"+username, http.StatusFound)
}

func products(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("htmlTemplates/startingPage.html", "htmlTemplates/header.html", "htmlTemplates/footer.html")
	t.Execute(w, nil)
	httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

}

func startingPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("htmlTemplates/startingPage.html", "htmlTemplates/header.html", "htmlTemplates/footer.html")
	fmt.Println(t)
	fmt.Println(err)
	Produkte := []Product{product1, product2}

	err = t.Execute(w, Produkte)

	fmt.Println(err)
}

func guthaben(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("htmlTemplates/guthaben.html", "htmlTemplates/header.html", "htmlTemplates/footer.html")
	currentlyLoggedInUser := GetActiveUser(r).Username
	// Weiterleitung auf formsCreditHandler erfolgt im html part !
	t.Execute(w, currentlyLoggedInUser)

	fmt.Println(err)

}

func contact(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("htmlTemplates/kontakt.html", "htmlTemplates/header.html", "htmlTemplates/footer.html")
	fmt.Println(err)
	t.Execute(w, nil)
}

func formsCredithandler(w http.ResponseWriter, r *http.Request) {

	loggedInUser := GetActiveUser(r)

	inventory, err2 := loadInventory(loggedInUser.Username)

	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Uups, something went wrong )-: 1 "))
		w.Write([]byte("Does the user, your tried to load your inventory from exists ?"))
		return
	}

	cobblestone := r.FormValue("valueCobblestone")
	cobbleStoneInt, err3 := strconv.Atoi(cobblestone)
	if err3 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Uups, something went wrong )-: 2 "))
		return
	}
	inventory.Cobblestone = inventory.Cobblestone + cobbleStoneInt
	inventory.save()

	// weiterleitung auf addCredits
	t, err := template.ParseFiles("htmlTemplates/addCredits.html")
	fmt.Println(err)

	err = t.Execute(w, inventory)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Uups, something went wrong )-: 3"))
		return
	}
	// valueCobblestone change, new account of cobblestone needs to be fixed

}

func registerUsers(w http.ResponseWriter, r *http.Request) {
	var response struct {
		Error string
		User  User
	}

	t, err := template.ParseFiles("htmlTemplates/registerUser.html")
	if r.Method == "GET" {
		// erster Request
		t.Execute(w, response)
		return
	}
	newUserName := r.FormValue("usernameRegister")

	if newUserName == "" {
		fmt.Println("Oh, die Variable scheint leer zu sein )-: ")
		response.Error = "Username can not be empty."
		t.Execute(w, response)
		return
	}
	fmt.Println(newUserName)
	if checkIfUserExists(newUserName) {
		fmt.Println("Oh, ein Nutzer mit einem solchen Namen scheint leider schon zu existieren")
		//w.Write([]byte("Username already exists"))
		response.Error = "Username already exists."
		t.Execute(w, response)
		return
	}
	fmt.Println("new user has been registered")

	newUser := User{
		Username: newUserName,
		Inv: Inventar{
			Username:    newUserName,
			Cobblestone: 0,
		},
	}
	// To prevent that an existing user is going to be overwritten only change, if userAlreadyExists = false

	newUser.save()

	fmt.Println(err)
	//err = t.Execute(w, newuser1)
	response.User = newUser
	err = t.Execute(w, response)
	if err != nil {
		fmt.Println(err)
	}

}

func loginUser(w http.ResponseWriter, r *http.Request) {
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
		fmt.Println(currentlyUser)
	}

	// try to log in user
	if checkIfUserExists(currentlyLoggedInUser) {
		Login(currentlyLoggedInUser)
		// testing because of logout in the html partcurrentlyLoggedInUserName login

		// Login without an valid username
	} else {
		//currentlyUser.errorWrongUsername =
		currentlyLoggedInUser = ""

	}
	t.Execute(w, currentlyUser.currentlyLoggedInUserName)
}

// error handeling here, later new file

func invalidUsername() error {
	return errors.New("this username is invalid, please register or log in with a valid one")
}

/*
 {{if ne .currentlyLoggedInUserName ""}}
                <h4 class="padding-top-small">Hallo {{.currentlyLoggedInUserName}}, du hast dich erfolgreich eingeloggt
                </h4>
                {{end}}
*/
