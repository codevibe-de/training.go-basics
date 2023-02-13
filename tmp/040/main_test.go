package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	s := Add(1, 2)
	if s != 3 {
		t.Errorf("Add(1, 2) = %d", s)
	}
}

func TestAdd_tableDriven(t *testing.T) {
	// declare test-data in tabular form
	testdata := []struct {
		a    int
		b    int
		want int
	}{
		{0, 0, 0},
		{1, 0, 1},
		{10, 20, 30},
		{-1, -1, -2},
	}
	// execute a range of tests using that data
	for _, data := range testdata {
		got := Add(data.a, data.b)
		if got != data.want {
			t.Fatalf("Add(%d, %d), got %d, want %d", data.a, data.b, got, data.want)
		}
	}
}

func TestAdd_tableDrivenSubTests(t *testing.T) {
	// declare test-data in tabular form
	testdata := []struct {
		a    int
		b    int
		want int
	}{
		{0, 0, 0},
		{1, 0, 1},
		{10, 20, 30},
		{-1, -1, -2},
	}
	// execute a range of tests using that data
	for _, data := range testdata {
		testname := fmt.Sprintf("Add(%d,%d)", data.a, data.b)
		t.Run(testname, func(t *testing.T) {
			got := Add(data.a, data.b)
			if got != data.want {
				t.Fatalf("%s, got %d, want %d", testname, got, data.want)
			}
		})
	}
}

func ExampleAdd() {
	fmt.Println(Add(-1, +1))
	// Output: 0
}
