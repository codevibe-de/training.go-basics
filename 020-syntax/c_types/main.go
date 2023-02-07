package main

import "fmt"

type centimeter int
type meter float32

func main() {
	fmt.Println("Types:")
	var lengthCentimetre centimeter = 140
	var lengthMeter meter = centimeterToMeter(lengthCentimetre)

	// Ausgabe: "140 cm equals 1.40 m"
	fmt.Printf("%d cm equals %.2f m\n", lengthCentimetre, lengthMeter)
}

func centimeterToMeter(cm centimeter) meter {
	return meter(cm) / 100.0
}
