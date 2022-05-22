//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//* Add three numbers together using any combination of the existing functions.
	//  * Print the result
	//* Call every function at least once
	fmt.Println(greet("north"))
	fmt.Println(add(1, 2, 3))
	fmt.Println(random())
	fmt.Println(randomTwos())
	fmt.Println(addFunc(random(), randomTwos))
}

func greet(name string) string {
	return fmt.Sprint("Hello", name, "!")
}

func add(x, y, z int) int {
	return x + y + z
}

func addFunc(x int, addFunc func() (int, int)) int {
	y, z := addFunc()
	return x + y + z
}

func random() int {
	return rand.Intn(100)
}

func randomTwos() (int, int) {
	return rand.Intn(100), rand.Intn(100)
}
