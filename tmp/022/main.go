package main

import (
	"crypto/md5"
	"fmt"
	"github.com/integrii/flaggy"
	"github.com/tauinger-de/training.go-basics.lib/compare"
	"golang.org/x/exp/slices"
)

func main() {
	var algFlag = "sha256"
	var inputArg = ""
	flaggy.SetName("gohash")
	flaggy.SetVersion("0.1")
	flaggy.String(&algFlag, "a", "alg", "The hashing algorithm to use")
	flaggy.AddPositionalValue(&inputArg, "input", 1, true, "the input to hash (file or URL)")
	flaggy.Parse()

	fmt.Println(inputArg)

	fmt.Println(compare.EqualInt(1, 2))

	md5.New()
	slices.Compare(nil, nil)
}
