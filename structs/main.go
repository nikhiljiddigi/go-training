package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstname string
	lastname  string
	contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson", contactInfo{"nikhil@gmail.com", 50090}}
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)
	// fmt.Println()
	// fmt.Println(alex.contactInfo)
	// fmt.Printf("%+v", alex.contactInfo)

	// var grey person
	// fmt.Println(grey)
	// fmt.Printf("%+v", grey)
	// grey.firstname = "Grey"
	// grey.lastname = "Anderson"
	// grey.contactInfo.email = "abc@gmail.com"
	// grey.contactInfo.zipCode = 642872
	// fmt.Println(grey)
	// grey.contactInfo.email = "xyz@gmail.com"
	// fmt.Printf("%+v", grey)

	brad := person{
		firstname: "Brad",
		lastname:  "Anderson",
		contactInfo: contactInfo{
			email:   "abc@gmail.com",
			zipCode: 0,
		},
	}
	brad.updateName("Brandon")              // Updating firstname using pointers
	brad = updateLastName(brad, "Peterson") // Updating lastname without using pointers
	brad.print()
}

func (pointerToPerson *person) updateName(newfirstName string) {
	(*pointerToPerson).firstname = newfirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func updateLastName(p person, newLastName string) person {
	p.lastname = newLastName
	return p
}
