package main

import "fmt"

func main() {
	poi := PointOfInterest{Coords{12, 7}, "Aral Tanke"}
	poi.x = 123
	fmt.Println(poi.Coords.String())
}

type Coords struct {
	x, y int
}

func (c Coords) String() string {
	return fmt.Sprintf("%d/%d", c.x, c.y)
}

type PointOfInterest struct {
	Coords
	name string
}
