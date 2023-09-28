package main

import (
	"fmt"
	"log"
	"net/http"
)

func hellohandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method Not Supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "helloworld!")
}

func formhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println(w, "parse form error", err)
		return
	}
	fmt.Fprintf(w, "POST Request Successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	class := r.FormValue("class")
	fmt.Fprintf(w, "Address = %s\n", address)
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "class = %s\n", class)
	fmt.Println(w, "Address = %s\n", address) //fprintf shows string values on the display webpage of the path
	// println will only print the value on console

}

func main() {
	fileserver := http.FileServer(http.Dir("./")) // all static files will also be served from here
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formhandler)
	http.HandleFunc("/hello", hellohandler)

	fmt.Println("Starting Server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
