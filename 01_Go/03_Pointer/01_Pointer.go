package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func main() {

	answer := 42
	fmt.Println(&answer) // Output: memory address of answer (e.g., 0xc0000100b0)
	
	address := &answer
	fmt.Println(*address) // Output: 42
	
	fmt.Printf("address is a %T\n", address) // Output: address is a *int

	// fmt.Println(*answer) // Dereferencing an int type directly would result in an error

	china := "China"

	var home *string
	fmt.Printf("home is a %T\n", home) // Output: home is a *string
	
	home = &china
	fmt.Println(*home) // Output: China

	// home = &answer // Error: cannot use &answer (type *int) as type *string in assignment

	// Declare a string pointer
	var administrator *string

	// Pointer points to the first person
	scolese := "Christopher J. Scolese"
	administrator = &scolese
	fmt.Println(*administrator) // Output: Christopher J. Scolese

	// Pointer now points to the second person
	bolden := "Charles F. Bolden"
	administrator = &bolden
	fmt.Println(*administrator) // Output: Charles F. Bolden

	// Modify the bolden variable, using the pointer to see the change
	bolden = "Charles Frank Bolden Jr."
	fmt.Println(*administrator) // Output: Charles Frank Bolden Jr.

	// You can also change the variable indirectly by "dereferencing"
	*administrator = "Maj. Gen. Charles Frank Bolden Jr."
	fmt.Println(bolden) // Output: Maj. Gen. Charles Frank Bolden Jr.

	// Assign the pointer to another variable, which will create a pointer to the same variable.
	major := administrator
	*major = "Maj. General Charles Frank Bolden Jr."
	fmt.Println(bolden) // Output: Maj. General Charles Frank Bolden Jr.

	fmt.Println(administrator == major) // Output: true

	lightfoot := "Robert M. Lightfoot Jr."
	administrator = &lightfoot // administrator now points to a new memory address
	fmt.Println(administrator == major) // Output: false 

	// However, dereferencing a pointer and assigning its value to another variable creates a copy.
	charles := *major
	*major = "Charles Bolden"
	fmt.Println(charles) // Output: Maj. General Charles Frank Bolden Jr.
	fmt.Println(bolden)  // Output: Charles Bolden

	// Even if two string variables point to different addresses, if their string values are the same, they are equal.
	charles = "Charles Bolden"
	fmt.Println(bolden == charles)   // Output: true; string comparison only requires the values to be equal.
	fmt.Println(&bolden == &charles) // Output: false; pointer comparison requires the memory addresses to be equal.

	// Pointer to a struct
	type person struct {
		name, superpower string
		age              int
	}
	
	timmy := &person{
		name: "Timothy",
		age:  10,
	}
	
	timmy.superpower = "flying"
	fmt.Printf("%+v\n", timmy) // Output: &{name:Timothy superpower:flying age:10}

	// Pointer to an array
	superpowers := &[3]string{"flight", "invisibility", "super strength"}

	fmt.Println(superpowers[0])   // Output: flight
	fmt.Println(superpowers[1:2]) // Output: [invisibility super strength]

}
