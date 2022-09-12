package main

import (
	"fmt"
	"log"
)

func main() {
	// string ranges can be accessed just as array

	someRandomStrings := "Testing strings ranges"
	fmt.Println(someRandomStrings[2:8])

	// convertings strings to number  use the convert package.
	// intToParse := "50"

	// fmt.Println(strconv.ParseInt(intToParse, 10, 30))
	log.Println("log msg")
}
