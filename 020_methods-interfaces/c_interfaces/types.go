package main

import "fmt"

// --- interfaces ---

type Formatter interface {
	format(n int) string
}

type Parser interface {
	parse(s string) int
}

// --- implementations ---

type BinaryFormatter struct {
}

func (bf BinaryFormatter) format(n int) string {
	return fmt.Sprintf("%b", n) // or strconv.FormatInt()
}

type BinaryParser struct {
}

func (bs BinaryParser) parse(s string) int {
	var result = 0
	_, _ = fmt.Sscanf(s, "%b", &result) // or strconv.ParseInt()
	return result
}

type LoggingParser struct {
	Parser
}

func (lp LoggingParser) parse(s string) int {
	fmt.Printf("About to parse: %q\n", s)
	return lp.Parser.parse(s)
}
