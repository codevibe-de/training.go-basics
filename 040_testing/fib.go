package main

var (
	prev     = 0
	prevPrev = 0
)

func Reset() {
	prev, prevPrev = 0, 0
}

func NextFib() int {
	res := prev + prevPrev
	prev, prevPrev = res, prevPrev
	return res
}
