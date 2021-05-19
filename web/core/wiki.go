package core

import (
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (page *Page) Save() error {
	filePath := "/tmp/" + page.Title + ".txt"
	return ioutil.WriteFile(filePath, page.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	filePath := "/tmp/" + title + ".txt"
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: content}, nil
}
