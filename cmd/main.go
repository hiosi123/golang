package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloWorldPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func handleFunction(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Hellow World!")
	case "/ninja":
		fmt.Fprint(w, "Wallace")
	default:
		fmt.Fprint(w, "big Fat Error!")
	}
	fmt.Printf("Handling function with %s request\n", r.Method)
}

func htmlVsPlain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("htmlVsPlain")
	w.Header().Set("Content-Type", "text/html") // implicitly set to html tag
	fmt.Fprint(w, "<h1>Hello World!</h1>")
	r.Body
}

func timeout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Timeout attempt")
	time.Sleep(2 * time.Second)
	fmt.Fprint(w, "Did not timeout!?")
}

func helloWorldNinjaMode(w http.ResponseWriter, r *http.Request) {
	fmt.Println("helloWorldNinjaMode")
	fmt.Fprint(w, "<h1 style=\"background-color:grey;\">Hello World</h1>")
}

func main() {
	// http.HandleFunc("/", htmlVsPlain)
	// http.HandleFunc("/timeout", timeout)
	// http.ListenAndServe("", nil)

	// server.ListenAndServe()
	server := http.Server{
		Addr:         "",
		Handler:      nil,
		ReadTimeout:  1000,
		WriteTimeout: 1000,
	}
	var muxNinjaMode http.ServeMux
	server.Handler = &muxNinjaMode
	muxNinjaMode.HandleFunc("/", helloWorldNinjaMode)

	server.ListenAndServe()
}

//https and http,
//your making the data secure, before sending the data to destinated server.
//domain. server ip address
