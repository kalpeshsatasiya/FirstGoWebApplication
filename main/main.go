package main

import (
	"net/http"
)

func main() {

	http.ListenAndServe(":8070", http.FileServer(http.Dir("public")) )
}
