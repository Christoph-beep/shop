package goFiles

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"text/template"
)

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
	// umbauen, nicht mehr über URL, sondern aus Form
	username := r.FormValue("userNameRegister")
	cobbleStoneString := r.FormValue("cobblestone")
	cobblestone, _ := strconv.Atoi(cobbleStoneString)
	p := Inventar{Username: username, Cobblestone: cobblestone}
	p.save()
	http.Redirect(w, r, "/view/"+username, http.StatusFound)
}

// for any mistakes
func error404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	t, _ := template.ParseFiles("htmlTemplates/404.html", "htmlTemplates/header.html", "htmlTemplates/footer.html")
	t.Execute(w, nil)

}

func products(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("htmlTemplates/startingPage.html", "htmlTemplates/header.html", "htmlTemplates/footer.html")
	t.Execute(w, nil)
	httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

	// Es gibt 2 wichtige Datentypen
	//http.Handler

	//ergebnis := http.HandlerFunc(error404)
	//ergebnis.ServeHTTP()

	//inv.ServeHTTP()

}

func (i Inventar) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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

// error
func contact(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("htmlTemplates/kontakt.html", "htmlTemplates/header.html", "htmlTemplates/footer.html")
	fmt.Println(err)
	t.Execute(w, nil)
}

func formsCredithandler(w http.ResponseWriter, r *http.Request) {

	loggedInUser := GetActiveUser(r)

	// User is not registered and therefore can not add credits
	/*if !checkIfUserExists(username) && username != "" {
		t, err := template.ParseFiles("htmlTemplates/guthaben.html")
		fmt.Println("user exists not l 98 ")
		fmt.Println(err)
		t.Execute(w, nil)
		return
	}*/
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
	IsUserLoggedIn := false
	type UserLoggIn struct {
		UserLoggedIn              bool
		currentlyLoggedInUserName string
	}

	t, err := template.ParseFiles("htmlTemplates/login.html")
	if err != nil {
		fmt.Println(err)
	}

	// If input = logout, the user wants to log out
	// logout information comes through the "name" fild in login

	currentlyLoggedInUser := r.FormValue("usernameLogin")
	usernameLogout := r.FormValue("logout1")

	currentlyUser := UserLoggIn{
		UserLoggedIn:              IsUserLoggedIn,
		currentlyLoggedInUserName: currentlyLoggedInUser,
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
		currentlyUser.UserLoggedIn = true

		// Login mit ungültigem Nutzernamen
	} else {
		error404(w, nil)
		currentlyLoggedInUser = ""
	}

}
