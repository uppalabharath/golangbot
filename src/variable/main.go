package main

import "fmt"

func main() {
	var age int // variable declaration var <var_name> <type>
	fmt.Println(age)
	age = 2
	fmt.Println(age)
	age = 4
	fmt.Println(age)

	//go does the type inference from the value assigned to the variable
	// when creating a variable with initial value we can omit the type
	// var name = "Test"
	var name string = "Test" // variable with an initial value

	fmt.Println(name)

	// initialising and decalaring multipe variable of same type
	var favNum, id = 10, 102

	// declaring two variables in a single statement
	var (
		address string
		zipcode int
	)

	//go doesnt allow unused variables so every variale declared shoulde be used somewhere
	address = "Street1"
	zipcode = 123456
	fmt.Println(favNum, id, address, zipcode)

	//shorthand declaration of variables

	int1, int2, str1 := 1, 2, "A"
	fmt.Println(int1, int2, str1)

	//shorthand syntax ":=" can be used only if atleast one of the variables is new
	// shorthand syntax always expects the variables to be initialized

}
