package main

import "fmt"

func sayGreetings(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBy(n string) {
	fmt.Printf("GoodBye %v \n", n)
}

func cycleGreetings(arrOfNames []string, f func(string)) {
	for _, value := range arrOfNames {
		f(value)
	}
}
func main() {
	// cycleGreetings([]string{"cedrick", "Junior", "prince", "Prosper"}, sayBy)
	var x int

	fmt.Println(&x)
}

// returning multiple values from a function.
//  add double  paranthesis as the func signature
func personDetails(n string, Dob int) (string, int, string) {
	return "Cedrick is", 40, "old"
}


// a function with zero or more parameters and a return value
func add(numbers ...int) int {
	result := 0
	for _, number := range numbers {
	result += number
	}
	return result
	}