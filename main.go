package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("this is my main main")

	mainWiki()
	// homeView()
	viewView()
	viewEdit()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
