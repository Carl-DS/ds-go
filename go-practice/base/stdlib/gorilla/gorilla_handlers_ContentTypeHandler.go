package gorilla

import (
	"github.com/gorilla/handlers"
	"io"
	"net/http"
)

func main() {
	http.Handle("/contentType", useContentTypeHandler(handler6()))
	http.ListenAndServe(":1234", nil)
}

func handler6() http.Handler {
	return http.HandlerFunc(myHandler6)
}

func myHandler6(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "text/plain")
	io.WriteString(rw, "Hello World")
}

func useContentTypeHandler(next http.Handler) http.Handler {
	return handlers.ContentTypeHandler(next, "application/x-www-form-urlencoded")
}
