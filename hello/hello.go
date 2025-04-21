package main

import (
	"fmt"
	"log"

	"go-learn/greetings"
)

func main() {
	// Set properties of the predefined variables
	// the log entry prefix and a flag to diable printing
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message
	message, err := greetings.Hello("Nick")

	// if an error was returned, print it to the console and exit
	if err != nil {
		log.Fatal(err)
	}

	// if no error was returned, print the message
	// to the console
	fmt.Println(message)

	// Array and Slice examples

	// Array use fix length
	array := [2]string{"Peeny", "Penny"}
	fmt.Println(array[0])
	fmt.Println(array[1])

	// Slice use dynamic length
	slice := []string{"Peeny", "Penny"}
	fmt.Println(slice[0])
	fmt.Println(slice[1])

	s := make([]byte, 5, 6) // when the length is 5, the capacity is 6. when it is inuse it will be 5
	fmt.Println(s)

	// s == []byte{0, 0, 0, 0, 0}
	fmt.Println(len(s)) // 5
	fmt.Println(cap(s)) // 6

	// slicing of slice and array

	x := []byte{'a', 'b', 'c', 'd', 'e'}

	// x[1:4] == []byte{'b', 'c', 'd'}, shearing the storage as x
	fmt.Println(x[0:2]) // [a b]
	fmt.Println(x[1:3]) // [b c]
	fmt.Println(x[2:4]) // [c d]

	y := [3]string{"sdf", "sd", "sdf"}
	z := y[:]      // [sdf sd sdf]
	fmt.Println(z) // [sdf sd sdf]
	// slice is a reference type, so it will share the same storage as the array. it is different from the array
	// that it is not completely copied. it is just a reference to the array
	// for example it does not copy the array, it just points to the same memory location

	d := []byte{'r', 'o', 'a', 'd'}
	e := d[2:] // [a d]
	// e == []byte{'a', 'd'}

	e[1] = 'm'
	// e == []byte{'a', 'm'}
	// d == []byte{'r', 'o', 'a', 'm'}  // d is changed because e is a reference to d
}
