package goFiles

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func calculate(zahl1 int, zahl2 int, operator string) int {
	if operator == "-" {
		ergebnis := zahl1 - zahl2
		return ergebnis
	}
	if operator == "+" {
		ergebnis := zahl1 + zahl2
		return ergebnis
	}
	if operator == "*" {
		ergebnis := zahl1 * zahl2
		return ergebnis
	}
	log.Panic("Operator is unfortunately not implemented")
	return 0
}

// term z.B 2 + 4

// term wird vollständig als String über die URL übergeben und muss daher einzeln betrachtet werden
func parseAndCalculate(term string) int {
	if strings.Contains(term, "+") {
		split := strings.Split(term, "+")
		zahl1, _ := strconv.Atoi(split[0])
		zahl2, _ := strconv.Atoi(split[1])
		return calculate(zahl1, zahl2, "+")
	}
	if strings.Contains(term, "-") {
		split := strings.Split(term, "-")
		zahl1, _ := strconv.Atoi(split[0])
		zahl2, _ := strconv.Atoi(split[1])
		return calculate(zahl1, zahl2, "-")
	}
	if strings.Contains(term, "*") {
		split := strings.Split(term, "*")
		zahl1, _ := strconv.Atoi(split[0])
		zahl2, _ := strconv.Atoi(split[1])
		return calculate(zahl1, zahl2, "*")
	}
	return 0
}

func mathhandler(w http.ResponseWriter, r *http.Request) {
	ergebnis := parseAndCalculate(r.URL.Path[len("/math/"):])
	fmt.Fprintln(w, "Ich bin ein Beispiel", ergebnis)
}
