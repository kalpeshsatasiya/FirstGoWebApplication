package main

import (
	"net/http"
	"strings"
	"os"
	"bufio"
)

func main() {
	http.Handle("/", new(MyHandler))

	http.ListenAndServe(":8070", nil)
}

type MyHandler struct {
	http.Handler
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := "public/" + req.URL.Path
	f, err := os.Open(string(path))

	if err == nil {

		bufferReader := bufio.NewReader(f)
		var contentType string
		if (strings.HasSuffix(path, ".css")) {
			contentType = "text/css"
		}else if (strings.HasSuffix(path, ".html")) {
			contentType = "text/html"
		}else if (strings.HasSuffix(path, ".js")) {
			contentType = "application/javascript"
		}else if (strings.HasSuffix(path, ".png")) {
			contentType = "image/png"
		}else if (strings.HasSuffix(path, ".jpg")) {
			contentType = "image/jpg"
		}else if (strings.HasSuffix(path, ".gif")) {
			contentType = "image/gif"
		}else if (strings.HasSuffix(path, ".mp4")) {
			contentType = "video/mp4"
		}else {
			contentType = "text/plain"
		}

		w.Header().Add("Content-Type", contentType)
		bufferReader.WriteTo(w)
	}else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}
}