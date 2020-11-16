package main

// Maps allow us to store things in a key, value system

import "fmt"

func main() {
	// two ways of defining a map named grades with string keys and float32 values
	// var grades map[string]float32
	grades := make(map[string]float32)

	grades["Timmy"] = 42
	grades["Jess"] = 92
	grades["Sam"] = 67

	fmt.Println(grades)

	TimsGrade := grades["Timmy"] // returns the value assigned to this key "Timmy"
	fmt.Println(TimsGrade)

	delete(grades, "Timmy") // deleting a key value pair from the grades map
	fmt.Println(grades)

	for k, v := range grades {
		fmt.Println(k, ":", v)
	}
}
