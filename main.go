package main

import (
	"log"
	"net/http"

	"github.com/semahmannaii/go-rest-api/configs"
	"github.com/semahmannaii/go-rest-api/controllers"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	db := configs.ConnectToDB()
	controller := controllers.Controller{}

	r := mux.NewRouter()

	r.HandleFunc("/", controller.GetMangas(db)).Methods("GET")
	r.HandleFunc("/{id}", controller.GetManga(db)).Methods("GET")
	r.HandleFunc("/", controller.CreateManga(db)).Methods("POST")
	r.HandleFunc("/", controller.UpdateManga(db)).Methods("PUT")
	r.HandleFunc("/{id}", controller.DeleteManga(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
