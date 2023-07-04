package main

import "fmt"

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

func main() {
	rect := rectangle{3, 4}
	fmt.Println(rect.area())
}
