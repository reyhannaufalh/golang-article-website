package main

import (
	"crud-product-golang/config"
	"crud-product-golang/controllers/articlecontroller"
	"crud-product-golang/controllers/categorycontroller"
	"crud-product-golang/controllers/homecontroller"
	"log"
	"net/http"
)

func main()  {
	config.ConnectDB()

	// Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/create", categorycontroller.Create)
	http.HandleFunc("/categories/store", categorycontroller.Store)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/update", categorycontroller.Update)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	// Articles
	http.HandleFunc("/articles", articlecontroller.Index)
	http.HandleFunc("/articles/create", articlecontroller.Create)
	http.HandleFunc("/articles/store", articlecontroller.Store)
	http.HandleFunc("/articles/edit", articlecontroller.Edit)
	http.HandleFunc("/articles/update", articlecontroller.Update)
	http.HandleFunc("/articles/delete", articlecontroller.Delete)

	log.Println("Server started on: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}