package goFiles

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

//// Important globally ////

// Check if user is logged in inventar l. 24 standardUser

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
// post request NICHT über die URl
// get über URL

////// 		 Authentification 			/////

// general functions

// saving users inventory
func saveHandler(w http.ResponseWriter, r *http.Request) {
	username := GetActiveUser(r).Inv.Username
	user, err := loadInventory(username)
	if err != nil {
		fmt.Println(err)
	}

	cobblestone := user.Cobblestone

	p := Inventar{Username: username, Cobblestone: cobblestone, Password: user.Password}
	p.saveInventar()

}

func loadcurrentlychoosenProduct(productID int) Product {
	choosenProductValue := loadProduct(productID)
	return choosenProductValue
}

func productsAvailable() []Product {
	ProductsAvailable := []Product{product0, product1, product2}
	return ProductsAvailable
}

func GetActiveUser(r *http.Request) User {
	return currentlyLoggedInUser
}

func Login(username string) error {
	user, err := loadUser(username)
	if err != nil {
		return err
	}
	currentlyLoggedInUser = user
	return nil
}

func Logout() {
	fmt.Println("ausloggen")
	currentlyLoggedInUser = User{}
}

func registerUsers(w http.ResponseWriter, r *http.Request) {
	var response struct {
		Error string
		Inventar
	}

	// end of general functions

	t, err := template.ParseFiles("htmlTemplates/registerUser.html", "htmlTemplates/header.html", "htmlTemplates/footer.html")
	if r.Method == "GET" {
		// first Request
		t.Execute(w, response)
		return
	}
	newUserName := r.FormValue("usernameRegister")
	newPasswordUsername := r.FormValue("usernamePasswordRegister")

	if newUserName == "" {
		fmt.Println("Oh, die Variable scheint leer zu sein )-: ")
		response.Error = "Username can not be empty."
		t.Execute(w, response)
		return
	}

	if newPasswordUsername == "" {
		fmt.Println("Oh, das Passwort scheint leer zu sein )-: ")
		response.Error = "Password can not be empty."
		t.Execute(w, response)
		return
	}

	fmt.Println(newUserName)
	// To prevent that an existing user is going to be overwritten only change, if userAlreadyExists = false
	if checkIfUserExists(newUserName) {
		fmt.Println("Oh, ein Nutzer mit einem solchen Namen scheint leider schon zu existieren")
		response.Error = "Username already exists."
		t.Execute(w, response)
		return
	}
	fmt.Println("new user has been registered")

	newUser := User{
		Inv: Inventar{
			Username:    newUserName,
			Cobblestone: 0,
			Password:    newPasswordUsername,
		},
	}

	// saving password
	newUser.save()

	fmt.Println(err)
	//err = t.Execute(w, newuser1)
	response.Username = newUserName
	err = t.Execute(w, response)
	if err != nil {
		fmt.Println(err)
	}

}

func startingPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("htmlTemplates/startingPage.html", "htmlTemplates/header.html", "htmlTemplates/footer.html")
	fmt.Println(err)
	// map over currently available products
	ProductsAvailableLocally := productsAvailable()

	err = t.Execute(w, ProductsAvailableLocally)

	fmt.Println(err)
}

func guthaben(w http.ResponseWriter, r *http.Request) {
	// Credits are added in the formsCreditHandler l.167
	t, err := template.ParseFiles("htmlTemplates/guthaben.html", "htmlTemplates/header.html", "htmlTemplates/footer.html")
	currentlyLoggedInUser := GetActiveUser(r).Inv.Username
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

	inventory, err2 := loadInventory(GetActiveUser(r).Inv.Username)

	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Uups, something went wrong )-: 1 "))
		w.Write([]byte("Does the user, your tried to load your inventory from exists ?"))
		return
	}

	cobblestone := r.FormValue("valueCobblestone")
	cobbleStoneInt, err3 := strconv.Atoi(cobblestone)
	inventory.Cobblestone = inventory.Cobblestone + cobbleStoneInt
	if err3 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Uups, something went wrong )-: 2 "))
		return
	}

	// weiterleitung auf addCredits
	t, err := template.ParseFiles("htmlTemplates/addCredits.html", "htmlTemplates/header.html", "htmlTemplates/footer.html")
	fmt.Println(err)

	err = t.Execute(w, inventory)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Uups, something went wrong )-: 3"))
		return
	}
	// valueCobblestone change, new account of cobblestone needs to be fixed

}

func purchaseSucessfully(w http.ResponseWriter, r *http.Request) {
	fmt.Println("purchaseSuccess")
	t, err := template.ParseFiles("htmlTemplates/purchaseSucessful", "htmlTemplates/header.html", "htmlTemplates/footer.html")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(err)

	actualLoggedInUserName := GetActiveUser(r).Inv.Username
	actualLoggedInUserCobblestone := GetActiveUser(r).Inv.Cobblestone
	inventory, err := loadInventory(actualLoggedInUserName)
	if err != nil {
		fmt.Println(err)
		fmt.Println("An error ocucured while trying to load the inventory l.216 handler")
	}

	QuantityChoosenProductsString := r.FormValue("QuantityChoosenProducts")

	productID_String := r.FormValue("SumbitProduct")
	productIDint, err := strconv.Atoi(productID_String)
	if err != nil {
		fmt.Println(err)
	}
	choosenProductValue := loadcurrentlychoosenProduct(productIDint)

	QuantityChoosenProductsInt, err := strconv.Atoi(QuantityChoosenProductsString)
	fmt.Println(QuantityChoosenProductsInt)
	fmt.Println("convertion was sucessful")
	if err != nil {
		fmt.Println("An error occured while trying to convert string to int l. 59 buying system")
		fmt.Println(err)
	}
	totalAmountToPay := QuantityChoosenProductsInt * int(choosenProductValue.Preis)

	// user decides to buy product, Creditcheck required
	if actualLoggedInUserCobblestone < totalAmountToPay {
		error404(w, r)
	}

	if err != nil {
		fmt.Println(err)
	}

	if QuantityChoosenProductsString == "" {
		QuantityChoosenProductsString = "0"
	}
	fmt.Println(QuantityChoosenProductsString)

	fmt.Println(totalAmountToPay)
	fmt.Println("Ist der zu zahlende Betrag")
	if err != nil {
		fmt.Println(err)
		fmt.Println("l. 245 handler")
	}

	inventory.Cobblestone = inventory.Cobblestone - totalAmountToPay
	inventory.saveInventar()

	t.Execute(w, inventory)
}

type userLogin struct {
	currentUser string
	errors      []internalErrors
}

type internalErrors struct {
	AuthentificationFailure string
}

var standardUser = userLogin{
	currentUser: currentlyLoggedInUser.Inv.Username,
}

func loginUser(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("htmlTemplates/login.html", "htmlTemplates/header.html", "htmlTemplates/footer.html")
	if err != nil {
		fmt.Println(err)
	}

	// If input = logout, the user wants to log out
	// logout information comes through the "name" fild in login

	///// Get information from the user /////
	currentlyLoggedInUser := r.FormValue("usernameLogin")
	currentpassword := r.FormValue("usernamePasswordLogin")
	// not a nice solution using a string called logout1 instead of a bool, but works currently
	usernameLogout := r.FormValue("logoutName")
	fmt.Println("Der aktuelle Wert aus usernameLogout ist " + usernameLogout)

	///// logout of user /////
	if usernameLogout == "logoutValue" {
		fmt.Println("logout")
		Logout()
		fmt.Println("Du wurdest erfolgreich ausgeloggt")
		t.Execute(w, GetActiveUser(r).Inv.Username)
		return
	}

	// try to log in user

	// add passwort auth with and
	if checkIfUserExists(currentlyLoggedInUser) {
		fmt.Println("user exists")

		if checkPassword(currentlyLoggedInUser, currentpassword) {
			Login(currentlyLoggedInUser)
			fmt.Println("user has been successfuly logged in l.287 handler")
			fmt.Println(GetActiveUser(r).Inv.Username + " ist der aktuelle Nutzername l.288 handler")
		}

		// testing because of logout in the html partcurrentlyLoggedInUserName login

		// Login without an valid username

	} else {
		authentificationFailure()
	}
	t.Execute(w, GetActiveUser(r).Inv.Username)

}

// error handeling here, later new file

func invalidUsername() error {
	return errors.New("this username is invalid, please register or log in with a valid one")
}

func authentificationFailure() error {
	return errors.New("Authentification was unfortunately not successful")
}

// errors

type RequestError struct {
	StatusCode int
	Err        error
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("status : err %v", r.Err)
}

func doRequest() error {
	return &RequestError{
		StatusCode: 503,
		Err:        errors.New("unavailable"),
	}
}
