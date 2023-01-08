package goFiles

import (
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

	r.HandleFunc("/save/{username}", saveHandler)
	r.HandleFunc("/startseite", startingPage)
	r.HandleFunc("/guthaben", guthaben)
	r.HandleFunc("/kontakt", contact)
	r.NotFoundHandler = http.HandlerFunc(error404)
	r.HandleFunc("/addCredits", formsCredithandler)
	r.HandleFunc("/registerUsers", registerUsers)
	r.HandleFunc("/login", loginUser)
	log.Fatal(http.ListenAndServe(":8080", r))
	// r.HandleFunc("/registerUserSuccess", registerUsersSuccess)

}

// handler
