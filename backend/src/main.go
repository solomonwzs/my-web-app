package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

const (
	_FrontendDir = "/home/solomon/workspace/javascript/my-web-app/dist"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func frontendHandler(w http.ResponseWriter, r *http.Request) {
	filename := filepath.Join(_FrontendDir, r.URL.Path[1:])
	fmt.Fprintln(w, filename)
}

func main() {
	fs := http.FileServer(http.Dir(_FrontendDir))
	http.Handle("/", fs)

	http.HandleFunc("/_f/", handler)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
