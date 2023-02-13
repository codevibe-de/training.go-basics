package main

var (
	prev     = 0
	prevPrev = 0
)

func Reset() {
	prev, prevPrev = 0, 0
}

// NextFib returns the next Fibonacci number
func NextFib() int {
	res := prev + prevPrev
	if res == 0 {
		res = 1
	}
	prev, prevPrev = res, prev
	return res
}
