package main

import (
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	frontendDir   = "/home/solomon/workspace/javascript/my-web-app/dist"
	strBaiduAppId = os.Getenv("BAIDU_APPID")
	baiduKey      = os.Getenv("BAIDU_KEY")
	baiduAppId    int64

	yeekitAppId = os.Getenv("YEEKIT_APPID")
	yeekitKey   = os.Getenv("YEEKIT_KEY")
)

func init() {
	rand.Seed(time.Now().UnixNano())
	baiduAppId, _ = strconv.ParseInt(strBaiduAppId, 10, 64)
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
	http.HandleFunc("/_backend/", proxyHandler)

	http.ListenAndServe(":8089", nil)
}
