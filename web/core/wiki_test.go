package core

import (
	"fmt"
	"testing"
)

func TestPage(t *testing.T) {
	page := &Page{Title: "test", Body: []byte("this is a test page")}
	err := page.save()
	if err != nil {
		t.Error("save err")
	}
	loadedPage, loadErr := loadPage("test")
	if loadErr != nil{
		t.Error("load error")
	}
	fmt.Println(string(loadedPage.Body))

}