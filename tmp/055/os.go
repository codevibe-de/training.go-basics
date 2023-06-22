package main

import (
	"bufio"
	"fmt"
	"io/fs"
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

func reduced() {
	filename := os.Args[1]
	outputDir, exists := os.LookupEnv("TEMP")

	os.Mkdir(outputDir, fs.ModeDir)

	var (
		dirs []os.DirEntry
		err  error
	)
	dirs, err = os.ReadDir("/var/tmp")

	var file *os.File
	file, err = os.Open(filename)

	var bytes []byte
	bytes, err = os.ReadFile(filename)

	os.Exit(0)

	fmt.Println(exists, dirs, err, file, bytes) // use vars
}
