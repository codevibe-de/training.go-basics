package main

import (
	"flag"
	"fmt"
	"os"
)

var globalInt int

func main() {
	var showUsage *bool = flag.Bool("usage", false, "Prints this usage message")
	var monsterName *string = flag.String("name", "", "Specifies the name of the monster")
	flag.IntVar(&globalInt, "age", 0, "Specifies the age of the monster in years")
	flag.Parse()
	if *showUsage || len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(0)
	}
	fmt.Println(*monsterName)
}
