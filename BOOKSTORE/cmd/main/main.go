package main

import (
	"GolangProjects/BOOKSTORE/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Starting server at port 9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
