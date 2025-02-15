package main

import "fmt"

func typeSwitch() {
	var x interface{} = 123
	switch i := x.(type) {
	case int:
		var someOtherInt int = i
		fmt.Println(someOtherInt)
	case bool, string:
		fmt.Println("type is bool or string")
	default:
		// ...
	}
}
