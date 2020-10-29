package main

import (
	"./routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/posts", routes.GetPosts).Methods("GET")
	router.HandleFunc("/posts", routes.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{id}", routes.GetPost).Methods("GET")
	router.HandleFunc("/posts/{id}", routes.UpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", routes.DeletePost).Methods("DELETE")

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}
