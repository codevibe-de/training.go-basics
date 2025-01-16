package main

import "fmt"

type Rectangle struct {
	width  int
	height int
}

func (r *Rectangle) area() int {
	return r.width * r.height
}

func (r *Rectangle) setWidth(w int) {
	(*r).width = w // muss man nicht so schreiben
}

func main() {
	rect := Rectangle{3, 4} // KEIN pointer
	fmt.Println(rect.area())
	rect.setWidth(10)
	fmt.Println(rect.area())
}

// 12
// 40
