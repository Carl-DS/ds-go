package gorilla

import (
	"compress/flate"
	"compress/gzip"
	"github.com/gorilla/handlers"
	"io"
	"net/http"
	"strings"
)

func main() {
	http.Handle("/gzip", useCompressLoggingHandler(handler5()))
	http.ListenAndServe(":1234", nil)
}

func handler5() http.Handler {
	return http.HandlerFunc(myHandler5)
}

func myHandler5(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "text/plain")
	io.WriteString(rw, "Hello World")
}

func useCompressLoggingHandler(next http.Handler) http.Handler {
	return handlers.CompressHandler(next)
}

func CompressHandlerLevel(h http.Handler, level int) http.Handler {
	if level < gzip.DefaultCompression || level > gzip.BestCompression {
		level = gzip.DefaultCompression
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	L:
		for _, enc := range strings.Split(r.Header.Get("Accept-Encoding"), ",") {
			switch strings.TrimSpace(enc) {
			case "gzip":
				w.Header().Set("Content-Encoding", "gzip")
				w.Header().Add("Vary", "Accept-Encoding")

				gw, _ := gzip.NewWriterLevel(w, level)
				defer gw.Close()

				h, hok := w.(http.Hijacker)
				if !hok {
					h = nil
				}

				f, fok := w.(http.Flusher)
				if !fok {
					f = nil
				}

				cn, cnok := w.(http.CloseNotifier)
				if !cnok {
					cn = nil
				}

				w = &compressResponseWriter{
					Writer:         gw,
					ResponseWriter: w,
					Hijacker:       h,
					Flusher:        f,
					CloseNotifier:  cn,
				}

				break L
			case "deflate":
				w.Header().Set("Content-Encoding", "deflate")
				w.Header().Add("Vary", "Accept-Encoding")

				fw, _ := flate.NewWriter(w, level)
				defer fw.Close()

				h, hok := w.(http.Hijacker)
				if !hok {
					h = nil
				}

				f, fok := w.(http.Flusher)
				if !fok {
					f = nil
				}

				cn, cnok := w.(http.CloseNotifier)
				if !cnok {
					cn = nil
				}

				w = &compressResponseWriter{
					Writer:         fw,
					ResponseWriter: w,
					Hijacker:       h,
					Flusher:        f,
					CloseNotifier:  cn,
				}

				break L
			}
		}

		h.ServeHTTP(w, r)
	})
}

type compressResponseWriter struct {
	io.Writer
	http.ResponseWriter
	http.Hijacker
	http.Flusher
	http.CloseNotifier
}

func (w *compressResponseWriter) WriteHeader(c int) {
	w.ResponseWriter.Header().Del("Content-Length")
	w.ResponseWriter.WriteHeader(c)
}

func (w *compressResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w *compressResponseWriter) Write(b []byte) (int, error) {
	h := w.ResponseWriter.Header()
	if h.Get("Content-Type") == "" {
		h.Set("Content-Type", http.DetectContentType(b))
	}
	h.Del("Content-Length")

	return w.Writer.Write(b)
}
