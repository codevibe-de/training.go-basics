package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// find all files ending with ".md"
	dirEntriesFound := make([]os.DirEntry, 0, 100)
	entries, _ := os.ReadDir(".")
	for _, e := range entries {
		if filepath.Ext(e.Name()) == ".md" {
			dirEntriesFound = append(dirEntriesFound, e)
		}
	}

	// show first three lines each by using a Scanner that splits at new-line chars
	for _, dirEntry := range dirEntriesFound {
		f, _ := os.Open(dirEntry.Name())
		fmt.Println(dirEntry.Name(), "\n------------------------")
		scanner := bufio.NewScanner(f)
		var n = 0
		for scanner.Scan() && n < 3 {
			fmt.Println(scanner.Text())
			n++
		}
	}
}
