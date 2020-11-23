package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Super Secret Information")
}

func handleRequests(){
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	fmt.Println("My Simple Server")
}