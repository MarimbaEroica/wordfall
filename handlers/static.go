package handlers

import "net/http"

func HandleStaticFiles() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
}
