package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	mapLab()
}

type product struct{}

func mapLab() {

	m := make(map[string]product)
	m["p-123"] = product{} // put

	// Eine Map hat eine Zahl von "Buckets", z.B. 10 Stück
	// jeder Bucket hat eine Liste an Elementen

	// PUT:
	// 1) berechne hash Wert vom Key -- "p-123" => 18731826761
	// 2) bucket-Index = hash % 10 = 1
	// 3) füge (Key,Item) der Liste in dem Bucket hinzu

	// 1) berechne hash Wert vom Key -- "p-4544" => 9879879871
	// 2) bucket-Index = hash % 10 = 1
	// 3) füge (Key,Item) der Liste in dem Bucket hinzu

	fmt.Println(m["P-123"]) // GET

	// GET:
	// 1) berechne hash Wert vom Key -- "p-123" => 18731826761
	// 2) bucket-Index = hash % 10
	// 3) habe nun einen Bucket mit 2 Einträgen...?
	// 4) Loop alle Einträge in Bucket, vergleiche Key, return bei Key-Match das Item

	jsonMap := make(map[string]any)
	jsonMap["firstName"] = "Thomas"
	jsonMap["zipCode"] = 97234
	jsonMap["favProduct"] = product{}
	bytes, _ := json.Marshal(jsonMap)
	fmt.Println(string(bytes))

	inp := `{
		"employees":[
			{"firstName":"John", "lastName":"Doe"},
			{"firstName":"Anna", "lastName":"Smith"},
			{"firstName":"Peter", "lastName":"Jones"}
		]
		}`

	// var intPtr *int = new(int)

	temp := make(map[string]any) // new == make + &
	res := &temp
	json.Unmarshal([]byte(inp), res)
	fmt.Println(*res)
}

func sliceLab() {
	arr := [3]string{"Guten", "Morgen", "Leute"}
	// slc := arr[:] // mehr Kapa
	slc := make([]string, 0, 10)
	slc = append(slc, (arr[:])...)
	fmt.Printf("%q\n", slc)

	arr[2] = "Ihr"
	fmt.Printf("%q\n", slc)

	slc = append(slc, "lieben", "Leute")

	// remove slice item at index 1
	slc = append(slc[:1], slc[2:]...)
	fmt.Printf("%q\n", slc)

	// insert item into slice at index 1
	slc = append(append(slc[0:1], "Abend"), slc[2:]...)
	fmt.Printf("%q\n", slc)
}
