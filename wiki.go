package main

import (
	"io/ioutil"
)

// Page -> some comment to get rid of linter msg.. or use lower case "page"
type Page struct {
	Title string
	Body  []byte
}

func mainWiki() {

	homePage1 := &Page{
		Title: "afk",
		Body:  []byte("Hello World from David's house!"),
	}
	homePage1.save()

}

// takes a receiver of p that points to "Page" data struct in RAM memory
func (p *Page) save() error {
	fileName := p.Title + ".txt"                    // p will be pointed to the Page struct ref
	return ioutil.WriteFile(fileName, p.Body, 0600) // 0600 file read-write (refer to linux permissions)
}

func loadPage(title string) (*Page, error) {
	fileName := title + ".txt"
	body, err := ioutil.ReadFile(fileName) // returns []byte (byte slice) data and error

	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
