package main

import (
	"fmt"
	"time"
)

func thisIsAFunction() {
}

func withParameters(n int) {
}

func returningSomething(name string, factor int) string {
	return ""
}

//func a_variables-and-constants() {
//	var n int
//	var n2 int = 0
//
//	var (
//		name string = "Enrico"
//		age  int    = 33
//	)
//
//	pi := 3.141592
//
//	city, zip := "Würzburg", 97080
//
//	streetWithNumber := "Bernerstr. 1"
//}

func constants() {
	const pi float32 = 3.141592
	const defaultName = "foo"

	const (
		kb = 1024
		mb = 1024 * 1024
	)
}

func foo(n int) {
	fmt.Printf("foo(%d)\n", n)
}

func iotaDemo() {
	const (
		zero = iota * 2
		two  = iota * 2
		four
		_
		eight
	)
	fmt.Println(zero, two, four, eight)
}

func pointer() {
	var strPointer *string
	text := "abc"
	strPointer = &text
	fmt.Println(strPointer, *strPointer)
}
func pointer2() {
	text := "abc"
	modify(&text)
	fmt.Println(text)
}
func modify(s *string) {
	*s = "new value"
}

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

func structs() {
	type address struct {
		zipCode int
		city    string
		country string
	}
	berlinCentral := address{zipCode: 10557, city: "Berlin"}
	fmt.Println(berlinCentral)

	hamburgFishmarket := address{22767, "Hamburg", "DE"}
	fmt.Println(hamburgFishmarket)
}

func structs2() {
	type address struct {
		zipCode int
		city    string
		country string
	}
	type contact struct {
		fullName    string
		homeAddress address
	}
	me := contact{"Thomas Auinger", address{97234, "Reichenberg", "DE"}}
	fmt.Println(me)
}

func anonStruct() {
	logData := struct {
		level     string
		message   string
		timestamp int64
	}{
		"INFO",
		"The application is starting up",
		time.Now().UnixMilli(),
	}
	fmt.Println(logData)
}

func taggedStruct() {
	type logData struct {
		level           string `json:"lvl" xml:"the-level"`
		message         string `json:"msg"`
		ignoreThisField int    `json:"-"`
	}
}

func withSameTypeParams(a, b, c int) {
}

func returningMultipleValues(name string) (int, error) {
	return 0, fmt.Errorf("Function not implemented")
}

func namedReturn(a, b, c, d int) (min, max int) {
	min, max = a, b
	return min, max
}

func printNames(intro string, names ...string) {
	fmt.Println(intro)
	fmt.Println(names)
}

func fibonacci(max int) (string, int) {
	return "", 0
}

func funcVars() {
	// define function type
	type filterFunc func(string) bool
	// implement it
	var isEmpty filterFunc = func(s string) bool {
		return len(s) == 0
	}
	// call it
	fmt.Println(isEmpty(""), isEmpty("abc"))
}

type filterFunc func(string) bool

func closure() {
	atLeastThree := createFilterFunc(3)
	fmt.Println(atLeastThree("ab"), atLeastThree("abc"))
}
func createFilterFunc(length int) filterFunc {
	return func(s string) bool {
		return len(s) >= length
	}
}

func main() {
	closure()
}
