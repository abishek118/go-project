package router

import (
	"github.com/abishek118/momgoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/items", controller.GetMyAllItems).Methods("GET")
	router.HandleFunc("/api/item", controller.CreateItem).Methods("POST")
	router.HandleFunc("/api/item/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/item/{id}", controller.DeleteAItem).Methods("DELETE")
	router.HandleFunc("/api/deleteallitem", controller.DeleteAllItems).Methods("DELETE")

	return router
}
