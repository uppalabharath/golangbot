/*
Defer is used to make the call execute just before the function which holds defer
1. defer can be used on functions as well as methods
2. defer calls maintain a stack and they are run in LIFO way.
3. Arguments for the defer are executed / evaluated when the function is executed not when the
function is called
*/
package main

import (
	"fmt"
	"sync"
)

type rect struct {
	length, width int
}

func (r rect) area(wg *sync.WaitGroup) {
	defer wg.Done()
	if r.length <= 0 {
		fmt.Printf("Rect %v length is invalid.\n", r)
		return
	}
	if r.width <= 0 {
		fmt.Printf("Rect %v width is invalid.\n", r)
		return
	}
	fmt.Printf("Rect %v area is %d.\n", r, (r.length * r.width))
}

func main() {
	name := "TestName"
	for _, v := range []rune(name) {
		defer fmt.Printf("%c\n", v)
	}
	var wg sync.WaitGroup
	rectangles := []rect{rect{10, 5}, rect{10, 10}, rect{-3, 4}, rect{4, -5}}

	for _, rectangle := range rectangles {
		wg.Add(1)
		rectangle.area(&wg)
	}
	a := 1
	// b holds the address of a
	// *b would dereference the address and gives the value of a
	b := &a
	fmt.Println(a, b, &a, *b, a == *b, &a == b)
}
