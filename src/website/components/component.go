package components

import (
	"fmt"
)

/*Author hold info about the post author*/
type Author struct {
	FirstName, LastName, Bio string
}

/*
Post holds the data of a blog post. */
type Post struct {
	title, content string
	postAuthor     Author
}

func (a Author) fullName() string {
	return fmt.Sprintf("%s, %s", a.FirstName, a.LastName)
}

func (p Post) details() {
	fmt.Println("Title:", p.title)
	fmt.Println("Content:", p.content)
	fmt.Println("Author:", p.postAuthor.fullName())
	fmt.Println("Bio:", p.postAuthor.Bio)
	fmt.Println("------------------------------------------------")
}

/*NewPost creates a new post*/
func NewPost(title, content string, author Author) Post {
	p := Post{title, content, author}
	return p
}

/*
Site is the site content
*/
type Site struct {
	Posts []Post
}

/*
Contents prints the blog posts with the author info
*/
func (s Site) Contents() {
	fmt.Println(":: Blog Contents ::")
	for _, post := range s.Posts {
		post.details()
	}
}
