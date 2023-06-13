package main

import (
	"fmt"
	"io"
)

type Stringer interface {
	String() string
}

func foo2(str Stringer) {
	fmt.Println(str.String())
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

type Base struct {
	b int
}

type Container struct {
	Base
	c string
}

func embedded() {
	cntr := Container{
		Base: Base{
			b: 1,
		},
	}
	fmt.Println(cntr.b == cntr.Base.b)
}

type ReadWriter interface {
	io.Reader
	io.Writer
}

type ReadWriteCloser interface {
	io.ReadCloser
	io.WriteCloser
}

type CountingReader struct {
	io.Reader
	BytesRead uint64
}

func (cr *CountingReader) Read(p []byte) (int, error) {
	n, err := cr.Reader.Read(p)
	cr.BytesRead += uint64(n)
	return n, err
}

func doSomeReading(r io.Reader) {
	bytes := make([]byte, 10)
	r.Read(bytes)
}

func main() {
	//cr := CountingReader{strings.NewReader("abc"), 0}
	//doSomeReading(&cr)
	//fmt.Println(cr.BytesRead)
	formatParseDemo()
	typeAssert()
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
