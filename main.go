package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", Login)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/refresh", Refresh)

	fmt.Println("Server is start...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
