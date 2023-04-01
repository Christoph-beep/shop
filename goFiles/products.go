package goFiles

type Product struct {
	Name     string
	Preis    float64
	ID       int
	BildName string
	Stock    int
}

// New Products

var product0 = Product{
	Name:     "WD-RED SSD 1TB",
	Preis:    110.50,
	ID:       0,
	BildName: "",
	Stock:    10,
}

var product1 = Product{
	Name:     "Seagate-Ironwolf 4TB",
	Preis:    92.47,
	ID:       1,
	BildName: "",
	Stock:    8,
}

var product2 = Product{
	Name:     "SanDisk Ultra 1TB ",
	Preis:    75.00,
	ID:       2,
	BildName: "",
	Stock:    12,
}

// function gives back requested product
func loadProduct(ProductID int) Product {
	//Produkte := []Product{product0, product1, product2}
	// look for a map insted of struct
	Produkte := make(map[int]Product)

	Produkte[0] = product0
	Produkte[1] = product1
	Produkte[2] = product2

	for i := range Produkte {
		if ProductID == Produkte[i].ID {
			// needs to be filled with specific values
			return Produkte[i]
		}

	}
	// if an error occures, an empty product is given back
	return Product{}
}

// interface for products under handler line 64
