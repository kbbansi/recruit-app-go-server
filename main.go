package main

import (
	"./pkg/config"
	"./pkg/logs"
	"fmt"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"time"
)

const port = "8297"
func main() {
	db, err:= config.DbConnect()
	HandleError(err)
	fmt.Println("Connected to Database")
	config.Database = db

	//Cross-Origin Scripting
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPost,
		http.MethodDelete,
		http.MethodOptions,
	})
	
	//Start API
	fmt.Println("Starting API on port", port)
	fmt.Println("API running on port", port)
	
	router:= NewRouter()
	
	server:= &http.Server{
		Addr:              fmt.Sprintf("127.0.0.1:%v", port),
		Handler:           handlers.CORS(origins, headers, methods)(router),
		WriteTimeout:       1 * time.Second,
		ReadTimeout:       1 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

func HandleError(err error) {
	if err != nil {
		logs.LogError(err)
		fmt.Println("Unable to setup resources")
		log.Panic(err)
	}
}