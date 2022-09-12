package main

import "fmt"

func sayGreetings(n string) {
	fmt.Printf("Good morning %v", n)
}

func main() {
	// greeting := "Hello good Morning"
	// fmt.Println(strings.Contains(greeting, "Hello"))

	// ages := []int{23, 45, 2, 3, 4, 5, 6, 7, 8}
	// sort.Ints(ages)
	// fmt.Println(ages)

	names := []string{"cedrick", "Junior", "prince", "Prosper"}

	for x := 0; x < len(names); x++ {
		// fmt.Println("value :", names[x])
	}

	// similar to for in
	for index, value := range names {
		fmt.Printf("the Value at index %v is %v \n", index, value)
	}

	// if er don't want to use the index value and vice versa is true
	for _, value := range names {
		fmt.Printf("the Value is %v \n", value)
	}
}
