package main

import "fmt"
import "net/http"
import "io/ioutil"
import "text/template"

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[6:]
	p, _ := loadPage(title)

	t, _ := template.ParseFiles("view.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil

}
