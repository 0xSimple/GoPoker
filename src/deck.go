package main

import "fmt"

type deck []string

func (d deck) print() { // First bracket is the receiver (Particular instance that can use the func)
	for i, card := range d {
		fmt.Println(i, card)
	}
}
