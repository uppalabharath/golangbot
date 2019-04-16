package main

import (
	"fmt"
)

type student struct {
	name, grade, country string
	age                  int
}

func filter(students []student, filterFunc func(s student) bool) []student {
	var filtered []student
	for _, student := range students {
		if filterFunc(student) {
			filtered = append(filtered, student)
		}
	}
	return filtered
}

func main() {
	students := []student{
		student{"A", "A", "India", 20},
		student{"A", "B", "India", 11},
		student{"A", "C", "India", 29},
		student{"A", "D", "USA", 30},
		student{"A", "E", "USA", 32},
	}

	ageFilter := func(student student) bool {
		return student.age > 15
	}

	filtered := filter(students, ageFilter)

	fmt.Println(filtered)
}
