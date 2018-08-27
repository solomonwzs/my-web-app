package main

import (
	"log"
	"net/http"
	"os"
)

var (
	frontendDir = os.Getenv("FRONTEND_DIST")
	listenAddr  = os.Getenv("BACKEND_ADDR")
)

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
	http.HandleFunc("/_backend/", proxyHandler)

	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
