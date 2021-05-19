package core

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body []byte
}

func (page *Page) save() error{
	filePath := "/tmp/" + page.Title + ".txt"
	return ioutil.WriteFile(filePath, page.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filePath := "/tmp/" + title + ".txt"
	content, err := ioutil.ReadFile(filePath)
	if err != nil{
		return nil, err
	}
	return &Page{Title: title, Body: content}, nil
}

func WebHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hi there, I love %s", r.URL.Path[1:])
}