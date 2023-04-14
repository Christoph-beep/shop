package goFiles

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Main() {

	// Gorilla default mux
	r := mux.NewRouter()

	// r.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// handler
	// Calculating productprice is in formscredithandler
	r.HandleFunc("/save/{username}", saveHandler)
	r.HandleFunc("/startseite", startingPage)
	r.HandleFunc("/guthaben", guthaben)
	r.HandleFunc("/kontakt", contact)
	r.NotFoundHandler = http.HandlerFunc(error404)
	r.HandleFunc("/addCredits", formsCredithandler)
	r.HandleFunc("/registerUsers", registerUsers)
	r.HandleFunc("/login", loginUser)
	// buyingSystem
	r.HandleFunc("/shoppingCart", buyingSystem)
	r.HandleFunc("/boughtProducts", productsAlreadyBoughtOverview)
	log.Fatal(http.ListenAndServe(":8080", r))
	r.HandleFunc("/purchaseSucessful", purchaseSucessfully)

	// testing

	r.Methods("GET").Path("/").HandlerFunc(endpointHandler)

}

func endpointHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint handler called")

	// r.HandleFunc("/registerUserSuccess", registerUsersSuccess)

}

// handler
