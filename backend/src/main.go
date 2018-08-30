package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

var (
	frontendDir = os.Getenv("FRONTEND_DIST")
	listenAddr  = os.Getenv("BACKEND_ADDR")
)

type fsHandler struct {
	http.Handler
	prevPath string
}

func (f *fsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	newUrl, _ := url.Parse(r.URL.Path[len(f.prevPath):])
	r.URL = newUrl
	f.Handler.ServeHTTP(w, r)
}

func setCORS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(
		"Access-Control-Request-Methods",
		"GET, POST, PUT")
	w.Header().Set(
		"Access-Control-Allow-Headers",
		r.Header.Get("Access-Control-Request-Headers"))
	w.Header().Set(
		"Access-Control-Allow-Origin",
		r.Header.Get("Origin"))
	w.Header().Set(
		"Access-Control-Allow-Credentials", "true")
}

func main() {
	fs := http.FileServer(http.Dir(frontendDir))

	http.Handle("/", fs)

	path := "/poorspeaker"
	http.Handle(path, &fsHandler{fs, path})

	path = "/about"
	http.Handle(path, &fsHandler{fs, path})

	http.HandleFunc("/_backend/", proxyHandler)

	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
