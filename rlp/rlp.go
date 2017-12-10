package main

import (
	"fmt"
)

// RLP (Recursive Length Prefix) is used by the Ethereum client to encode arbitrarily nested
// arrays of binary data.

// Encode currently will only take a string
// In future I will add support for arbitrarily nested strings
func Encode(input interface{}) ([]byte, error) {
	switch input.(type) {
	case string:
		s := []rune(input.(string))
		if len(s) == 1 && s[0] < 128 {
			return []byte(input.(string)), nil
		}
		encLen, err := encodeLength(len(s), 128)
		return append(encLen, []byte(input.(string))...), err
	default:
		return nil, nil
	}
}

func encodeLength(length int, offset byte) ([]byte, error) {
	if length < 56 {
		return []byte{byte(length) + offset}, nil // would this byte array always be 1 byte?
	}
	// should really return error if length >= int(math.Pow(256, 8))
	asBin := intToBigEndian(length)
	return []byte{byte(len(asBin)) + offset + 55}, nil
}

func intToBigEndian(len int) []byte {
	if len == 0 {
		return []byte(nil) // appending to nil slice just ignores the nil slice
	}
	return append(intToBigEndian(int(len/256)), byte(len%256))
}

func main() {
	thing, _ := Encode("dog")
	fmt.Println(thing)

	thing, _ = Encode("Đ")
	fmt.Println(thing)

	thing, _ = Encode("Lorem ipsum dolor sit amet, consectetur adipisicing elit")
	fmt.Println(thing)
}
