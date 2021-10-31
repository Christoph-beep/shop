package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Profile struct {
	Age   int
	Name  string
	Email string
	Money float64
}
type Product struct {
	Id          int
	Price       float64
	Name        string
	Description string
}

var karl_heinz Profile = Profile{
	Name:  "Karl Heinz",
	Age:   22,
	Money: 239.12,
	Email: "karl.heinz@shop.de",
}

var hdd_wd_red Product = Product{
	Id:          0,
	Name:        "Western Digital Red",
	Price:       92.32,
	Description: "Hoch performante HDD, die sich für den 24/7 Dauerbetrieb bestens eignet",
}

var minecraftRang Product = Product{
	Id:          1,
	Name:        "Minecraft-Gold-Rang",
	Price:       25,
	Description: "Rang, der dir eine golden glänzende Rüstung verleiht",
}

var all_products map[int]Product = make(map[int]Product)

// 0 -> hdd_wd_red
// 1 -> minecraftRang

func main() {
	var alter int = 42
	fmt.Printf("Ich bin %d Jahre alt %s\n", alter, "YUHUUUU!!!")
	fmt.Fprintf(os.Stdout, "Ich bin %d Jahre alt %s\n", alter, "YUHUUUU!!!")

	load_products()

	r := mux.NewRouter()

	r.HandleFunc("/profile", render_profile)
	r.HandleFunc("/products", render_products)
	r.HandleFunc("/product/{id:[0-9]+}", render_product)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8123", nil))
}

func load_products() {
	all_products[0] = hdd_wd_red
	all_products[1] = minecraftRang
}

func render_profile(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./html_templates/profile.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, karl_heinz)
	if err != nil {
		panic(err)
	}
}

func render_products(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./html_templates/products.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func render_product(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_as_string := vars["id"]
	fmt.Printf("Trying to render id : %s\n", id_as_string)
	product_id, _ := strconv.Atoi(id_as_string)
	tmpl, err := template.ParseFiles("./html_templates/product.html")
	if err != nil {
		fmt.Printf("Error while Handling Request %s: %s", r.RequestURI, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		monkey, _ := template.ParseFiles("./html_templates/monkey.html")
		monkey.Execute(w, nil)
		return
	}
	var product Product
	product, ok := all_products[product_id]
	if !ok {
		// throw_404
		render_404(w, r)
		return
	}
	_ = tmpl.Execute(w, product)
}

func render_404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	page_404, _ := template.ParseFiles("./html_templates/404.html")
	page_404.Execute(w, nil)
}
