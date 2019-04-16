/*
Explains the concept of inheritance in Golang
1. Any type that defines all the methods of a interface is said to be implementing the interface
2. A variable of type interface can hold all the type of its implememtors
*/
package main

import (
	"fmt"
)

/*Income is an interface holding the income info a project*/
type Income interface {
	source() string
	netIncome() int
}

/*FixedBidProject is the project where the income is fixed*/
type FixedBidProject struct {
	projectName string
	bidAmount   int
}

/*TimeAndMaterialProject is the project where the income varies as per the resource conbsumption*/
type TimeAndMaterialProject struct {
	projectName   string
	hourlyRate    int
	resourceCount int
}

func (proj FixedBidProject) source() string {
	return proj.projectName
}

func (proj TimeAndMaterialProject) source() string {
	return proj.projectName
}

func (proj FixedBidProject) netIncome() int {
	return proj.bidAmount
}

func (proj TimeAndMaterialProject) netIncome() int {
	return (proj.hourlyRate * proj.resourceCount)
}

func main() {
	income1 := FixedBidProject{"proj1", 100}
	income2 := FixedBidProject{"proj2", 100}
	income3 := TimeAndMaterialProject{"proj3", 100, 3}

	incomes := []Income{income1, income2, income3}
	sum := 0
	for _, project := range incomes {
		fmt.Printf("Project %s generated %d dollars. \n", project.source(), project.netIncome())
		sum += project.netIncome()
	}
	fmt.Printf("Total income generated: %d dollars. \n", sum)
}
