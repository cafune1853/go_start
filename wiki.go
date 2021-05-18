package go_start

type Page struct {
	Title string
	Body []byte
}

func (page *Page) save() error{
	filePath := "/tmp/" + page.Title + ".txt"

}