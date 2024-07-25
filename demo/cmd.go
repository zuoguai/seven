package main

import (
	"log"
	"net/http"
	"seven/zuoWeb/zuo"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	web := zuo.New()
	web.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello "))
	})
	web.GET("/zuo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Zuo"))
	})
	web.GET("/guai", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Guai"))
	})

	log.Fatal(http.ListenAndServe(":18080", web))

}

// type Zuo struct{}

// func (zuo *Zuo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		w.Write([]byte("Hello Guai"))
// 	case "/zuo":
// 		w.Write([]byte("Hello Zuo"))
// 	default:
// 		w.Write([]byte("404"))
// 	}
// }
