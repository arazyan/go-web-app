package main

import (
	"fmt"
	"net/http"
)

func handleFunction(w http.ResponseWriter, r *http.Request) {
	// request routing
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Hello World")
	case "/whoami":
		fmt.Fprint(w, "just a sad archlinux user")
	default:
		fmt.Fprint(w, "%%%Big Fat Error!%%%")
	}
	// figuring out what type of request we are doing
	// it is GET, wow
	fmt.Printf("Handling function with %s request\n", r.Method)
}

func htmlText(w http.ResponseWriter, r *http.Request) {
	// by default implicitly was set to "text/html"
	// w.Header().Set("Content-Type", "text/html")

	// manually set it to plain text
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, "<h1>Hello World</h1>")
}

func main() {
	// http.HandleFunc() отвечает за mapping путей и страниц
	// http.HandleFunc("/", handleFunction)
	http.HandleFunc("/", htmlText)

	// 	err := http.ListenAndServe(":8080", nil)
	// 	if err != nil {
	// 		fmt.Println(err)
	// }

	// configure a server
	server := http.Server{
		Addr:    "",
		Handler: nil,
		//
		ReadTimeout: 1000, // 1000ms=1s
	}

}
