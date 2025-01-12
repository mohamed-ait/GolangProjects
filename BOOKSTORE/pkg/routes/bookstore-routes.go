package routes

import (
<<<<<<< HEAD
	"GolangProjects/BOOKSTORE/pkg/controllers"

=======
>>>>>>> 958d93c3ff041de2048b563e5d6ce1868817fd6d
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
<<<<<<< HEAD
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{booId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{booId}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{booId}", controllers.UpdateBook).Methods("PUT")

=======
	router.HandleFunc("/book/", controllers.createBook).Methods("POST")
>>>>>>> 958d93c3ff041de2048b563e5d6ce1868817fd6d
}
