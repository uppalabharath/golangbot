/*
This class explains the use of structs same as classes in Golang
*/
package main

import "oop/employee"

func main() {
	e := employee.New("Test", "Abcd", 10, 5)
	e.LeavesRemaining()
	// Zero value struct doesnt have a meaning
	// So need to implement new functionm to create proper objects
	// If a package defines a type T then the convention is to name it newT(params)
	// If the package has only one type then the convention is new(param)
}
