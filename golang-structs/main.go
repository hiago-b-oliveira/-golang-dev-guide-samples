package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	// contactInfo contactInfo // alternative sintax
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func main() {
	// alex := person{"Alex", "Anderson"}
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	var john person
	fmt.Println(john) // zero values
	fmt.Printf("%+v\n", john)
	john.firstName = "John"

	marry := person{
		firstName: "Marry",
		lastName:  "Jane",
		contactInfo: contactInfo{
			email:   "marry@gmail.com",
			zipCode: 1000,
		},
	}
	marry.updateName("M")
	marry.print()

	/* ******************** */
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
	}
	fmt.Println(colors)

	var colors2 map[string]string
	fmt.Println(colors2)

	colors3 := make(map[string]string)
	colors3["white"] = "#ffffff"
	delete(colors3, "white")

	for k, v := range colors {
		fmt.Println(k, v, "Aqui")
	}
}
