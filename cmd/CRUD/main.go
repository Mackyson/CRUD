package main

import (
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func WriteName(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	io.WriteString(w, params.ByName("name"))
}
func main() {
	router := httprouter.New()
	router.GET("/:name", WriteName)
	log.Println("Listen on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
