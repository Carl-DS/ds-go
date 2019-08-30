package gorilla

import (
	"github.com/gorilla/handlers"
	"io"
	"net/http"
)

func main() {
	http.Handle("/analysis", useCanonicalHost(handler7()))
	http.ListenAndServe(":1234", nil)
}

func handler7() http.Handler {
	return http.HandlerFunc(myHandler7)
}

func myHandler7(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "text/plain")
	io.WriteString(rw, "Hello World")
}

func useCanonicalHost(next http.Handler) http.Handler {
	return handlers.CanonicalHost("http://test.gs-robot.com/", http.StatusFound)(next)
}
