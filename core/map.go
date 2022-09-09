package main

import "fmt"

func main() {
	// name map<keys type> <valuse types>
	menu := map[string]float64{
		"Bread": 4.99,
		"pie":   4.89,
		"salad": 8.99,
		"pizza": 9.99,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])
	// looping maps

	for key, value := range menu {
		fmt.Println(key, "-", value)
	}

}
