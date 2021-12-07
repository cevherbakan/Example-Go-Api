package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"

)

func main() {
	router := mux.NewRouter()



	fmt.Println("Listening at 8080")
	http.ListenAndServe(":8080",router)
}

func Home(w http.ResponseWriter, r *http.Request){
	router.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hell")
	})
}