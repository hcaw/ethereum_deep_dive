package main

import (
	"fmt"
	// "reflect"
)

// RLP (Recursive Length Prefix) is used by the Ethereum client to encode arbitrarily nested
// arrays of binary data.

func main() {
	thing, _ := Encode("hello")
	fmt.Println(thing)
}

// Encode currently will only take an arbitrarily nested arrays of strings
// In future I will add support for more types
func Encode(input interface{}) ([]byte, error) {
	switch input.(type) {
	case string:
		return []byte{0, 1}, nil
	default:
		return []byte{0}, nil
	}
}
