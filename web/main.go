package main

import (
	"fmt"
	"github.com/cafune1853/go_start/web/core"
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
	page, _ := core.LoadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", page.Title, page.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("edit Handler invoked")
	title := r.URL.Path[len("/edit/"):]
	page, err := core.LoadPage(title)
	if err != nil{
		page = &core.Page{Title: title}
	}
	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</textarea><br>"+
		"<input type=\"submit\" value=\"Save\">"+
		"</form>", page.Title, page.Title, page.Body)
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


