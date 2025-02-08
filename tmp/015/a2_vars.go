package main

import "fmt"

func vars_and_constants() {
	var n int
	var n2 = 0

	var (
		name string = "Enrico"
		age         = 33
	)

	pi := 3.141592

	city, zip := "Würzburg", 97080

	street := "Bernerstr. 1"

	fmt.Println(n, n2, name, age, pi, city, zip, street)
}
