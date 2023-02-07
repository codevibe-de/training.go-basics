package main

import (
	"fmt"
	"math"
	"strings"
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
	return 0, fmt.Errorf("function not implemented")
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

func deferDemo() {
	defer fmt.Println("you're gonna be fine")
	fmt.Println("Sooner or")
	defer func() {
		fmt.Println("later")
	}()
}

func stringsDemo() {
	s := "Hey you"
	strings.Split(s, " ")
}

func arrayDemo() {
	var arr [3]int
	fmt.Println(arr)
	arr[1] = 8888
	fmt.Println(arr)

	strArr := [3]string{"Hey you", "Out there in the cold"}
	fmt.Printf("%q (%T)\n", strArr, strArr)
}

func forDemo() {
	var squares [5]int
	for i := 0; i < 5; i++ {
		squares[i] = i * i
	}
	fmt.Println(squares)
}

func forDemo2() {
	x := 100.0
	for x > 2 {
		x = math.Sqrt(x)
		fmt.Printf("%.3f ", x)
	}
}

func forDemoRange() {
	factors := [4]string{"cero", "uno", "dos", "tres"}
	for i, v := range factors {
		fmt.Printf("#%d %s, ", i, v)
	}
}

type rectangle struct {
	width  int
	height int
}

func (r *rectangle) area() int {
	return r.width * r.height
}

func (r *rectangle) setWidth(w int) {
	(*r).width = w
}

func sliceDemo() {
	intSlice := []int{1, 2, 3}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 4)
	fmt.Println(intSlice)
}

func sliceLenAndCap() {
	intSlice := []int{1, 2, 3}
	intSlice = append(intSlice, 4)
	fmt.Printf("length: %d, capacity: %d", len(intSlice), cap(intSlice))
}

func slicing() {
	boolArray := [4]bool{true, true, false, false}
	boolSlice := boolArray[1:3]
	fmt.Println(boolSlice)
}

func makeSlice() {
	var strSlice []string = make([]string, 2, 10)
	fmt.Printf("%q\n", strSlice)
	fmt.Printf("length: %d, capacity: %d", len(strSlice), cap(strSlice))
}

func switchDemo() {
	n := 43
	switch mod := n % 2; mod {
	case 0:
		fmt.Printf("%d is an even number\n", n)
	default:
		fmt.Printf("%d is NOT an even number\n", n)
	}
}

func main() {
	switchDemo()
	//r := rectangle{3, 4}
	//fmt.Println(r.area())
	//r.setWidth(10)
	//fmt.Println(r.area())
}

func ifDemo() {
	if res, err := connect(); err == nil {
		handleError(err)
	} else {
		doStuffWithResult(res)
	}
}

func doStuffWithResult(res string) {

}

func handleError(err error) {

}

func connect() (string, error) {
	return "", nil
}
