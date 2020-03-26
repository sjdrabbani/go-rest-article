package main

import (
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title string 'json:"title"'
	Desc string 'json:"desc"'
	Content string 'json:"content"'
}

func homePage(w http.ResponseWriter, r *http.Request) {	
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
