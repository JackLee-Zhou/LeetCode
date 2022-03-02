package main

import (
	"fmt"
	"testing"
)

func TestMulti(t *testing.T) {
	fmt.Println(Multi("498828660196", "840477629533"))
}

func TestPermute(t *testing.T) {
	fmt.Println(Permute([]int{0, 1}))
	fmt.Println(Permute([]int{1, 2, 3}))
}
