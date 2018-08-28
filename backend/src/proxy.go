package main

import (
	"io"
	"net/http"
	"strconv"
)

const (
	userAgent = "Mozilla/5.0 (Linux; Android 5.1.1; Nexus 6 Build/LYZ28E) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Mobile Safari/537.36"
)

func proxyRequest(w http.ResponseWriter, r *http.Request) {
	forward := r.Header.Get("X-Forward")
	if forward == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req, err := http.NewRequest(r.Method, forward, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for k, v := range r.Header {
		if k == "Content-Length" {
			l, _ := strconv.Atoi(v[0])
			req.ContentLength = int64(l)
		} else if k != "X-Forward" {
			req.Header.Set(k, v[0])
		}
	}
	req.Header.Set("User-Agent", userAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		w.Header().Set(k, v[0])
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	setCORS(w, r)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	proxyRequest(w, r)
}
