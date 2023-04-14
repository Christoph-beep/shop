package goFiles

import (
	"encoding/json"
	"fmt"
	"os"
)

// Cobblestone zum bezahlen
// Nutzername zum identifizieren
type Inventar struct {
	Username    string `json:"username"`
	Cobblestone int    `json:"cobble"`
	Password    string `json:"password"`
}

// registering of users

var currentlyLoggedInUser User

type User struct {
	// Username string
	Inv            Inventar
	isUserLoggedIn bool
}

func (u User) save() error {
	return u.Inv.saveInventar()
}

// Writing part , saveInventar definition
// with this the Inventar is going to be saved
func (i Inventar) saveInventar() error {
	fmt.Println("User is going to be saved")
	// Empty file is created
	filename := "Users/" + i.Username + ".txt"
	user, err := json.Marshal(Inventar{Username: i.Username, Cobblestone: i.Cobblestone, Password: i.Password})

	if err != nil {
		fmt.Println("error occured")
	}
	fmt.Println(string(user))
	// Converting string to int needs to be done in the reading part
	//return os.WriteFile(filename, []byte(strconv.Itoa(i.Cobblestone)), 0600)
	return os.WriteFile(filename, []byte(user), 0600)

}

// Reading part
func checkPassword(username string, password string) bool {
	localUser, err := loadInventory(username)
	if err != nil {
		return false
	}
	if password == localUser.Password {
		return true
	}
	return false
}

func loadInventory(username string) (Inventar, error) {
	filename := "Users/" + username + ".txt"
	inventory := Inventar{}

	umzugskarton, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error occured")
		return Inventar{}, err
	}

	// unmarshalling the json data
	err = json.Unmarshal(umzugskarton, &inventory)

	if err != nil {
		fmt.Println("Leider ist ein Fehler beim laden des Inventars aufgetreten", username, err)
		return Inventar{}, err
	}

	return inventory, nil
}

func loadUser(username string) (User, error) {
	Inventar, err := loadInventory(username)

	if err != nil {
		return User{}, err
	}

	return User{
		Inv: Inventar,
	}, nil
}

// checks

func checkIfUserExists(username string) bool {
	filename := "Users/" + username + ".txt"
	_, checkValue := os.Stat(filename)

	if os.IsNotExist(checkValue) {
		fmt.Printf("%v file exists\n", checkValue)
		fmt.Println("false")
		return false

	} else {
		fmt.Printf("%v file exist\n", checkValue)
		fmt.Println("true")

		return true
	}

}
