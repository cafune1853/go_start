package main

import (
	"github.com/cafune1853/go_start/web/core"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", core.WebHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


