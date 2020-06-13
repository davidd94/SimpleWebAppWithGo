package main

import (
	"fmt"
	"log"
	"net/http"
)

// func homeHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hi there, I love %s", r.URL.Path[1:]) // reads http request pointer data struct
// }

// func homeView() {
// 	http.HandleFunc("/", homeHandler)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func viewView() {
	http.HandleFunc("/view/", viewHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
