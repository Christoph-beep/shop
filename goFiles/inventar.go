package goFiles

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var currentlyLoggedInUser User

// Cobblestone zum bezahlen
// Nutzername zum identifizieren
type Inventar struct {
	Username    string
	Cobblestone int
}

// registering of users

type User struct {
	Username string
	Inv      Inventar
}

func (u User) save() error {
	return u.Inv.save()
}

func (i Inventar) save() error {
	filename := "Users/" + i.Username + ".txt"
	return os.WriteFile(filename, []byte(strconv.Itoa(i.Cobblestone)), 0600)
}

func loadInventory(username string) (Inventar, error) {
	filename := "Users/" + username + ".txt"
	cobbleStoneAnzahl, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Leider ist ein Fehler beim laden des Inventars aufgetreten", username, err)
		return Inventar{}, err
	}
	cobbleStoneString := string(cobbleStoneAnzahl)
	cobbleStoneInt, _ := strconv.Atoi(cobbleStoneString)

	return Inventar{
		Username:    username,
		Cobblestone: cobbleStoneInt,
	}, nil
}

func loadUser(username string) (User, error) {
	Inventar, err := loadInventory(username)

	if err != nil {
		return User{}, err
	}

	return User{
		Username: username,
		Inv:      Inventar,
	}, nil
}

// checks

func checkIfUserExists(username string) bool {
	filename := "Users/" + username + ".txt"
	_, checkValue := os.Stat(filename)

	if os.IsNotExist(checkValue) {
		fmt.Printf("%v file does not exist\n", checkValue)
		return false
	} else {
		fmt.Printf("%v file exist\n", checkValue)
		return true
	}
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
