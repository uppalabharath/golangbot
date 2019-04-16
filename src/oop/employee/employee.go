package employee

import "fmt"

/*
Employee Struct to hold employee data
*/
type Employee struct {
	firstName     string
	lastName      string
	totalLeaves   int
	leavesAvailed int
}

/*
LeavesRemaining Calculates the leaves remaining
*/
func (e Employee) LeavesRemaining() {
	fmt.Printf("Leaves remaining for Employee: %s, %s are %d\n", e.firstName, e.lastName, e.totalLeaves-e.leavesAvailed)
}

/*
New creates a new employee
*/
func New(first, last string, total, availed int) Employee {
	e := Employee{first, last, total, availed}
	return e
}
