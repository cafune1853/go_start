package main

import (
	"errors"
	"fmt"
	"github.com/cafune1853/go_start/web/core"
	"html/template"
	"log"
	"net/http"
	"regexp"
)

var ValidPattern = regexp.MustCompile("^/(view|edit|save)/([a-zA-Z0-9]+)$")

func getTitle(path string) (string, error){
	matched := ValidPattern.FindStringSubmatch(path)
	if matched == nil{
		return "", errors.New("wrong path")
	}
	return matched[2], nil
}

func makeHandler(cf func(w http.ResponseWriter, r *http.Request, title string)) http.HandlerFunc{
	return func(writer http.ResponseWriter, request *http.Request) {
		title, err :=getTitle(request.URL.Path)
		if err != nil {
			http.NotFound(writer, request)
			return
		}
		cf(writer, request, title)
	}
}

func simpleHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("simple Handler invoked")
	fmt.Fprintf(w, "Hi there, I love %s", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string){
	fmt.Println("view Handler invoked")
	page, err := core.LoadPage(title)
	if err != nil{
		renderPage(w, "not_found.html", &core.Page{Title: title})
		return
	}
	renderPage(w, "view.html", page)
}

func renderPage(w http.ResponseWriter, templateFile string, page *core.Page) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	t.Execute(w, page)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string){
	fmt.Println("edit Handler invoked")
	page, err := core.LoadPage(title)
	if err != nil{
		page = &core.Page{Title: title}
	}
	renderPage(w, "edit.html", page)
}

func saveHandler(w http.ResponseWriter, r *http.Request,title string){
	fmt.Println("save Handler invoked")
	content := r.FormValue("body")
	pageToSave := &core.Page{Title: title, Body: []byte(content)}
	pageToSave.Save()
	http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

func main(){
	http.HandleFunc("/", simpleHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}


