package main

import (
	"fmt"
	"github.com/urfave/negroni"
	"io"
	"net/http"
)

func main() {
	n := negroni.New()
	n.UseFunc(printAuthorInfo)

	router := http.NewServeMux()
	router.Handle("/", handler5())

	n.UseHandler(router)
	n.Run(":3456")
}

func printAuthorInfo(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("www")
	fmt.Println("lds")
	next(w, r)
}

func handler5() http.Handler {
	return http.HandlerFunc(myHandler5)
}

func myHandler5(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, "Hello World")
}
