package routes

import (
	"github.com/gorilla/mux"
	"github.com/jrose8775/golang-playground/go/src/go-bookstore/pkg/controllers"
)

var RegisterBookStoresRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookid}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookid}", controllers.UpdateBookById).Methods("PUT")
	router.HandleFunc("/book/{bookid}", controllers.deleteBookById).Methods("DELETE")
}
