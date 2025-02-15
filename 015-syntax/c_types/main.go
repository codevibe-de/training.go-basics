package main

import "fmt"

type centimeter int
type meter float32

func main() {
	fmt.Println("Types:")
	var lengthCentimeter centimeter = 140
	var lengthMeter meter = centimeterToMeter(lengthCentimeter)

	// Ausgabe: "140 cm equals 1.40 m"
	fmt.Printf("%d cm equals %.2f m\n", lengthCentimeter, lengthMeter)
}

func centimeterToMeter(cm centimeter) meter {
	return meter(cm) / 100.0
}
