package main

import (
	"io"
	"log"
	"net/http"

	"CRUD/internal/handler"
	"CRUD/internal/storage"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func WriteName(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	io.WriteString(w, params.ByName("name"))
}
func main() {

	storage.Migrate()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	router := httprouter.New()
	router.GET("/:name", WriteName)
	router.POST("/signup", handler.CreateNewUser)
	router.DELETE("/delete/:name", handler.DeleteUser)
	handler := c.Handler(router) //CORSオプションの設定
	log.Println("Listen on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
