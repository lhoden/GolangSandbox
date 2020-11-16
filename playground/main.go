package main

import (
	"fmt"
	"math"
	"math/rand"
)

func foo() {
	fmt.Println("I just called foo. Pretty cool, eh?")
}

/* only need one float64 type defined because both x and y are float32 */
func add(x, y float64) float64 {
	return x + y
}

/* the "string" outside of parenthesis defines what is being returned */

func singular(a string) string {
	return a
}

/* You can return multiple things! But you have to define all of the types you are returning */
func multiple(a, b string) (string, string) {
	return a, b
}

func main() {
	fmt.Println("Nothing", math.Sqrt(25))
	foo()
	fmt.Println("A number from 1-100", rand.Intn(100))

	// defining for the add method call
	/* This is the same as var num1, num2 float32 = 5.6, 9.5 */
	num1, num2 := 5.6, 9.5
	fmt.Println(add(num1, num2))

	w1, w2 := "Hey", "there"
	fmt.Println(multiple(w1, w2))
	//This will just return both of the words as two strings separated by a space

	var a int = 62
	var b float64 = float64(a) // converting an int to a type float64
	x := a                     // x will be type int
	// Returns the numbers we just created here
	fmt.Println("Testing something", a, b, x)

	/* Pointers */
	c := 15
	d := &c // the memory address
	fmt.Println(d)
	// If you want to point to something, use &
	// If you want to read the value stored in that location, use *
	fmt.Println(*d)

	// Creating a loop in Go
	i := 0
	for i < 10 {
		fmt.Println()
		i += 5
	}

	for x := 5; x < 25; x += 3 {
		fmt.Println("Do stuff", x)
	}
}
