package main

import (
	"html/template"
	"net/http"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func viewView() {
	http.HandleFunc("/view/", viewHandler)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save() // will return nil if success? and will return an error if error occurs
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func viewEdit() {
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles("./templates/" + tmpl + ".html")
	if err != nil {
		// http.Error function sends a specified HTTP response code (in this case "Internal Server Error") and error message.
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p) // will return nil if success? and will return an error if error occurs
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
