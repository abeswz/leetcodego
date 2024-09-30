package main

import "fmt"

var hashindex int32

// Get index of value by byte
func index(b rune) byte {
	return byte(b) - 'a'
}

func add(value rune) {
	v := leftShifting(value)
	hashindex |= v
}

func exist(value rune) bool {
	v := leftShifting(value)
	return hashindex&v != 0
}

func leftShifting(value rune) int32 {
	return int32(1 << index(value))
}

func main() {
	add('a')
	add('b')
	add('z')

	fmt.Println(exist('a'))
	fmt.Println(exist('b'))
	fmt.Println(exist('z'))
	fmt.Println(exist('t'))
}
