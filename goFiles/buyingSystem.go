package goFiles

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func buyingSystem(w http.ResponseWriter, r *http.Request) {

	type errorMessages struct {
		ErrorMessageNoUserLoggedIn   string
		ErrorMessageNotEnoughCredits string
	}

	type productInfo struct {
		choosenProduct      Product
		choosenProductName  string
		choosenProductPrice int
		errorMessages
	}

	ProductValue := r.FormValue("SumbitProduct")
	productIDint, err := strconv.Atoi(ProductValue)
	if err != nil {
		fmt.Println(err)
	}
	// gives back the product
	choosenProductValue := loadProduct(productIDint)

	var productInfo1 = productInfo{
		choosenProduct:      choosenProductValue,
		choosenProductName:  choosenProductValue.Name,
		choosenProductPrice: int(choosenProductValue.Preis),
		//errorMessages:       errorMessages{ErrorMessageNoProduct: "No Product has been choosen"},
	}

	fmt.Println("line 33 buying System")

	t, err := template.ParseFiles("htmlTemplates/shoppingCart.html", "htmlTemplates/header.html", "htmlTemplates/footer.html")
	if err != nil {
		fmt.Println(err)
	}

	// for purchasing a product, user needs to be logged in
	if GetActiveUser(r).Inv.Username != "" {
		// a user seems to be logged in
		fmt.Println("user is logged in")
		if choosenProductValue.Name != "" {
			fmt.Println(choosenProductValue.Name + " ist der Name deines ausgewählten Produkts")

			if err != nil {
				fmt.Println(ProductValue + "das steht als Productvalue")
				fmt.Println("An error occured during trying to convert string to int, likely the productID is empty")
				error404(w, r)
			}
			// if all went right, this is executed
			t.Execute(w, productInfo1.choosenProductName)

			// user decides to buy product
			QuantityChoosenProductsString := r.FormValue("QuantityChoosenProducts")
			QuantityChoosenProductsInt, err := strconv.Atoi(QuantityChoosenProductsString)
			if err != nil {
				fmt.Println("An error occured while trying to convert string to int l. 59 buying system")
				fmt.Println(err)
			}

			// proof if user has enough credits
			actualCredits := GetActiveUser(r).Inv.Cobblestone
			totalAmountToPay := QuantityChoosenProductsInt * int(productInfo1.choosenProduct.Preis)
			if actualCredits > totalAmountToPay {
				fmt.Println("not enough credits available")
				// Fehlercode über struct ?
			}

			// show the total amount 2 pay
		}
	}

	// no user seems to be logged in
	if GetActiveUser(r).Inv.Username == "" {
		fmt.Println("no user seems to be logged in")
		t.Execute(w, productInfo1.choosenProductName)
	}

}
