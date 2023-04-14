package goFiles

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func buyingSystem(w http.ResponseWriter, r *http.Request) {

	// user who is currently logged in
	currentlyActiveUser := GetActiveUser(r).Inv.Username

	// gives back the product, Funtion in general Funcions
	productID_String := r.FormValue("SumbitProduct")
	fmt.Println(productID_String + "ist die Nummer des Produkts l.17 buying System")
	productIDint, err := strconv.Atoi(productID_String)
	fmt.Println(productIDint)
	fmt.Println("convertion was successful")
	if err != nil {
		fmt.Println(err)
		fmt.Println("error l.20 buyingSystem")
	}
	choosenProductValue := loadcurrentlychoosenProduct(productIDint)

	// fmt.Println("line 33 buying System")
	t, err0 := template.ParseFiles("htmlTemplates/shoppingCart.html", "htmlTemplates/header.html", "htmlTemplates/footer.html")
	if err != nil {
		fmt.Println(err)
		fmt.Println("error l.31 buyingSystem")
	}

	// for purchasing a product, user needs to be logged in
	if currentlyActiveUser != "" {
		// a user seems to be logged in
		fmt.Println("user is logged in")
		if choosenProductValue.Name != "" {
			fmt.Println(choosenProductValue.Name + " ist der Name deines ausgewählten Produkts")
			// Checking if user has put Products into the shoppingCart
			// formsCredithandler(w, r)

			if err0 != nil {
				fmt.Println("An error occured during trying to convert string to int, likely the productID is empty")
				error404(w, r)
			}
			// if all went right, this is executed
			t.Execute(w, choosenProductValue)

			// user decides to buy product

		}
	}

	// no user seems to be logged in
	if GetActiveUser(r).Inv.Username == "" {
		fmt.Println("no user seems to be logged in")
		http.Redirect(w, r, "login", http.StatusMovedPermanently)
	}

}

func productsAlreadyBoughtOverview(w http.ResponseWriter, r *http.Request) {
	// Overview of the already bought items

}

func succesfullyPurchasedProducts(w http.ResponseWriter, r *http.Request) {

}
