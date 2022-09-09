package main

// blue print types
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bill
func newBills(name string) bill {

	// defining a new bill object
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}
