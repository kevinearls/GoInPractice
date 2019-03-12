package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	fmt.Println("Starting on port " + port)
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":" + port, nil)

}

func homePage(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(response, request)
		return
	}
	fmt.Fprint(response, "This is the homepage\n")
}
