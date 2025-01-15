package main

import (
	"fmt"
	"io"
)

type Stringer interface {
	String() string
}

func foo(s Stringer) {
	fmt.Println(s.String())
}

type SomeThing struct{}

func (st SomeThing) String() string {
	return "foo"
}

var stringer Stringer = SomeThing{}

var x interface{} = 3
var y interface{} = true
var z = struct{}{}
var xx any = x

type AbstractEntity struct {
	id int
}

func (b AbstractEntity) GetId() int {
	return b.id
}

type Product struct {
	AbstractEntity
	c string
}

func embedded() {
	product := Product{
		AbstractEntity: AbstractEntity{
			id: 1,
		},
	}
	product.GetId()
}

// true

type ReadWriter interface {
	io.Reader
	io.Writer
}

type ReadWriteCloser interface {
	io.ReadCloser
	io.WriteCloser
}

type Formatter interface {
	format(n int) string
}
type Parser interface {
	parse(s string) int
}

type BinaryFormatter struct {
}

func (bf BinaryFormatter) format(n int) string {
	return fmt.Sprintf("%b", n)
}

type BinaryParser struct {
}

func (bs BinaryParser) parse(s string) int {
	var result = 0
	fmt.Sscanf(s, "%b", &result)
	return result
}

func formatParseDemo() {
	f := BinaryFormatter{}
	s := f.format(42)
	fmt.Println(s)
	p := LoggingParser{BinaryParser{}}
	n := p.parse(s)
	fmt.Println(n)
}

type LoggingParser struct {
	Parser
}

func (lp LoggingParser) parse(s string) int {
	fmt.Printf("About to parse: %q\n", s)
	return lp.Parser.parse(s)
}

func typeAssert() {
	var n interface{} = int32(123)
	i, ok := n.(int32)
	if ok {
		fmt.Println("Variable `n` is actually of type `int32` with value", i)
	}
}
