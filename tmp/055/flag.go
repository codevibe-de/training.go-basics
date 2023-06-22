package main

import (
	"flag"
	"fmt"
	"os"
)

var age int

func main() {
	var showUsage *bool = flag.Bool("usage", false, "Show this usage message")
	var name *string = flag.String("name", "", "Your name")
	flag.IntVar(&age, "age", 0, "Your age")
	flag.Parse()
	if *showUsage || len(os.Args) <= 1 {
		flag.Usage()
		os.Exit(0)
	}
	fmt.Println(*name, age)
}
