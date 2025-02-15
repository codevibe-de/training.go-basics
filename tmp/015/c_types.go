package main

import "fmt"

func customTypes() {
	type äpfel int
	type birnen int

	var anzahlÄpfel äpfel = 40
	var anzahlBirnen birnen = 2

	// Invalid operation: anzahlÄpfel > anzahlBirnen (mismatched types äpfel and birnen)

	// okay mit Typ-Umwandlung
	mehrÄpfelAlsBirnen := int(anzahlÄpfel) > int(anzahlBirnen)
	fmt.Println(mehrÄpfelAlsBirnen)
}
