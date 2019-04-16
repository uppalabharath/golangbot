package main

import (
	"fmt"
	"github.com/uppalabharath/greeting"
	greetML "github.com/uppalabharath/greeting/v2"
)

func main() {
	fmt.Println(greet.Greet("Bharath"))
	greeting, err := greetML.Greet("Bharath", "fr")
	fmt.Println(greeting, err)
}
