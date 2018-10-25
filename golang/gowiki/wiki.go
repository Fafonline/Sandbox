package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

//Page struct
type Page struct {
	Title string
	//slice
	Body []byte
}

//Function on object

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	//init body as slice
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Cannot read file %s, error:%s", filename, err.Error())
		return nil, err
	}
	//Allocate new structure
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl)
	t.Execute(w, p)
}
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view.html", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit.html", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	body := r.FormValue("body")
	title := r.URL.Path[len("/save/"):]
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	//allocate and assign struct
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page")}
	//call received function
	p1.save()
	p2, err := loadPage("TestPage")
	if err != nil {
		fmt.Println("Cannot load file")
	}
	fmt.Println(string(p2.Body))

	appPort := "8080"

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	fmt.Printf("Listening on port %s\n", appPort)
	log.Fatal(http.ListenAndServe(":"+appPort, nil))
}
