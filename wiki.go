package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func mainWiki() {

	homePage1 := &Page{
		Title: "David's Homepage",
		Body:  []byte("Hello World from David's house!"),
	}
	homePage1.save()

	homePage2, _ := loadPage("David's Homepage")
	// homePage2, _ := loadPage(homePage1.Title)
	fmt.Println(string(homePage2.Body))

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
