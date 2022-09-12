package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a New Biil: ", reader)
	b := newBill(name)
	fmt.Println("New Bill Created:-", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose Option (a -  add Item, s - save bill, t - Add tip): ", reader)
	switch opt {
	case "a":
		name, _ := getInput("Item Name: ", reader)
		price, _ := getInput("Item Pric: ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The Price Must be a number")
			promptOptions(b)
		}
		b.addItemsToBill(name, p)
		fmt.Println("Item Added", name, price)
		promptOptions(b)

	case "t":
		tip, _ := getInput("Enter Tip Amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip Must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Tip added - ", tip)
		promptOptions(b)
	case "s":
		b.saveBill()
		fmt.Println("you save  bill", b.name)
	default:
		fmt.Println("that was not a valid option!...")
		promptOptions(b)
	}
}

func main() {
	myBill := createBill()
	promptOptions(myBill)

}
