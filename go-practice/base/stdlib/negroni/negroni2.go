package main

import (
	"github.com/urfave/negroni"
	"io"
	"net/http"
)

func main() {
	n := negroni.Classic()

	mux := http.NewServeMux()
	mux.Handle("/", handler2())
	mux.HandleFunc("/lds", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "www.lds.com")
		io.WriteString(w, "wechat")
	})

	n.UseHandler(mux)
	n.Run(":2345")
}

func handler2() http.Handler {
	return http.HandlerFunc(myHandler2)
}

func myHandler2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, "Hello World")
}
