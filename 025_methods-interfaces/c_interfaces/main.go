package main

import "fmt"

func main() {
	fmt.Println("Interfaces:")
	check(35, BinaryFormatter{}, BinaryParser{})
	check(35, BinaryFormatter{}, LoggingParser{BinaryParser{}})
}

func check(n int, f Formatter, p Parser) {
	out := p.parse(f.format(n))
	if out != n {
		panic(fmt.Sprintf("got %d, wanted %d", out, n))
	}
}
