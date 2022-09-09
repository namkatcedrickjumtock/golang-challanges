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
	cycleGreetings([]string{"cedrick", "Junior", "prince", "Prosper"}, sayBy)
}

// returning multiple values from a function.
//  add double  paranthesis as the func signature
func personDetails(n string, Dob int) (string, int, string) {
	return "Cedrick is", 40, "old"
}
