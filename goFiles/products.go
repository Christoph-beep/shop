package goFiles

type Product struct {
	Name     string
	Preis    float64
	ID       int
	BildName string
}

// New Products
var product1 = Product{
	Name:     "Seagate-Ironwolf 4TB",
	Preis:    92.47,
	ID:       0,
	BildName: "",
}

var product2 = Product{
	Name:     "WD-RED SSD 1TB",
	Preis:    110.50,
	ID:       1,
	BildName: "",
}
