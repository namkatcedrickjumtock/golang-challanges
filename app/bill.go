package main

import (
	"fmt"
	"os"
)

// blue print types
// a struct is a blue print which describes a type of data
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bill
func newBill(name string) bill {

	// defining a new bill object
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// receiver functinos

func (b *bill) formate() string {
	fs := "Bill Break Down \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ....$%v \n", k+":", v)
		total += v
	}

	// add tip
	fs += fmt.Sprintf("%-25v ...$%v\n", "Tip:", b.tip)

	fs += fmt.Sprintf("%-25v ...$%0.2f", "Total:", total+b.tip)
	return fs
}

// update the tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add item to bill
// adding pointers to receiver functions so it points to the original struct since Go createa copy of the struct each time it's being reference
func (b *bill) addItemsToBill(name string, price float64) {
	b.items[name] = price
}

func (bill *bill) saveBill() {
	data := []byte(bill.formate())

	err := os.WriteFile("bills/"+bill.name+".txt", data, 0644)

	if err != nil {
		panic(err)

	}
}
