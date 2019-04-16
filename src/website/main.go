/*
This file explains the concept of composition in Golang
*/

package main

import (
	"website/components"
)

func main() {
	author1 := components.Author{
		FirstName: "A1",
		LastName:  "B1",
		Bio:       "Tech Blogger1"}
	author2 := components.Author{
		FirstName: "A2",
		LastName:  "B2",
		Bio:       "Tech Blogger2"}
	author3 := components.Author{
		FirstName: "A3",
		LastName:  "B3",
		Bio:       "Tech Blogger3"}
	post1 := components.NewPost("Title 1", "Content 1", author1)
	post2 := components.NewPost("Title 2", "Content 3", author2)
	post3 := components.NewPost("Title 3", "Content 3", author3)

	site1 := components.Site{Posts: []components.Post{post1, post2, post3}}
	site1.Contents()
}
