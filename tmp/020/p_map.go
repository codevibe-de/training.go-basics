package main

import "fmt"

func main() {
	m := make(map[int]string) // type is `map[int]string`
	m[1] = "one"

	s2 := m[2]         // returns "" as default zero-value for string
	s1, exists := m[1] // returns "one", true

	l := len(m) // 1
	delete(m, 1)
	_, exists = m[1] // false

	fmt.Println(s1, s2, exists, l)
}

func shortForm() {
	zipToCityMap := map[string]string{
		"97074": "Würzburg",
		"97234": "Reichenberg",
	}
	fmt.Println(zipToCityMap)
	// Output: map[97074:Würzburg 97234:Reichenberg]
}

func complexKey() {
	type fancyMapKey struct {
		primaryId   int
		secondaryId string
	}
	fancyMap := map[fancyMapKey]string{
		fancyMapKey{1, "2"}: "test",
	}
	fmt.Println(fancyMap)
	// map[{1 2}:test]
}
