package main

import (
	"fmt"
	"github.com/cafune1853/go_start/web/core"
	"html/template"
	"log"
	"net/http"
)

func simpleHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("simple Handler invoked")
	fmt.Fprintf(w, "Hi there, I love %s", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("view Handler invoked")
	title := r.URL.Path[len("/view/"):]
	page, err := core.LoadPage(title)
	if err != nil{
		renderPage(w, "not_found.html", &core.Page{Title: title})
	}else{
		renderPage(w, "view.html", page)
	}
}

func renderPage(w http.ResponseWriter, templateFile string, page *core.Page) {
	t, _ := template.ParseFiles(templateFile)
	t.Execute(w, page)
}

func editHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("edit Handler invoked")
	title := r.URL.Path[len("/edit/"):]
	page, err := core.LoadPage(title)
	if err != nil{
		page = &core.Page{Title: title}
	}
	renderPage(w, "edit.html", page)
}

func saveHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("save Handler invoked")
	title := r.URL.Path[len("/save/"):]
	content := r.FormValue("body")
	pageToSave := &core.Page{Title: title, Body: []byte(content)}
	pageToSave.Save()
	http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

func main(){
	http.HandleFunc("/", simpleHandler)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


