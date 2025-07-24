package rout

import (
	"github.com/gorilla/mux"
	"github.com/vimukthiwaththegama/gorm-rest/pkg/controller"
)

var RegisterBookStoreRoutes = func(r *mux.Router){
	r.HandleFunc("/book/create",controller.CreateBook).Methods("POST")
	r.HandleFunc("/book/{id}", controller.GetBookById).Methods("GET")
	r.HandleFunc("/book/update/{id}", controller.UpdateBookById).Methods("PUT")
	r.HandleFunc("/book/",controller.GetAllBooks).Methods("GET")
	r.HandleFunc("/book/delete/{id}", controller.DeleteBookById).Methods("DELETE")
}