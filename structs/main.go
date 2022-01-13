package main

import "fmt"

type contactInfo struct {
	address string
	zipCode int
}

type person struct {
	lastName  string
	firstName string
	contactInfo
}

func main() {
	rakesh := person{
		lastName:  "Androtula",
		firstName: "Rakesh",
		contactInfo: contactInfo{
			address: "Hyderabad",
			zipCode: 500050,
		},
	}

	rakesh.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// func print(p person){
// 	fmt.Printf("%+v", p)
// }