package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("this is my main main")

	mainWiki()
	viewMain()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
