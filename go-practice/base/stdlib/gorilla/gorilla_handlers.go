package gorilla

import (
	"io"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	http.Handle("/", useLoggingHandler(handler()))
	http.Handle("/combined", useCombinedLoggingHandler(handler()))
	http.ListenAndServe(":1234", nil)
}

func handler() http.Handler {
	return http.HandlerFunc(myHandler4)
}

func myHandler4(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	io.WriteString(rw, "hello world")
}

func useLoggingHandler(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, next)
}

func useCombinedLoggingHandler(next http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, next)
}
