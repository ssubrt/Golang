package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to mod in golang");
	greeter()
	r := mux.NewRouter();
	r.HandleFunc("/",serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080",r));
}

func greeter(){
	fmt.Println("hey there from greeter");
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang </h1>"))
}