package main

import (
	_"fmt"
	"net/http"
	"log"
)


func test(w http.ResponseWriter, r *http.Request) {
	
}


func main(){

	router := http.NewServeMux()
	router.HandleFunc("/test", test)

	log.Println("Listning")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}

