package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	routes "github.com/hananx07/go-bookstore/PKG/Routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost :9010", r))
}
