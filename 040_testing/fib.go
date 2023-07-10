package main

var (
	prev     int
	prevPrev int
)

func Reset() {
	prev, prevPrev = 0, 1
}

func NextFib() int {
	res := prev + prevPrev
	prev, prevPrev = res, prev
	return res
}
