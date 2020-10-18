package main

import (
	"io"
	"log"
	"net/http"

	"CRUD/internal/handler"
	"CRUD/internal/storage"

	"github.com/julienschmidt/httprouter"
)

func WriteName(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	io.WriteString(w, params.ByName("name"))
}
func main() {

	storage.Migrate()

	router := httprouter.New()
	router.GET("/:name", WriteName)
	router.POST("/signup", handler.CreateNewUser)
	router.DELETE("/delete", handler.DeleteUser)
	log.Println("Listen on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
