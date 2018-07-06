package main

import (
	"fmt"
	"testing"
)

func TestSqrt(t *testing.T) {
	const in, out = 2, 4
	if x := Sqrt(in); x != out {
		t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
	}
}
func ExampleHello() {
	fmt.Println("hello1")
	// Output: hello
}
