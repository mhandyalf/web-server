package main

import (
	"fmt"
	"net/http"
	"web-server/handlers"
)

func main() {
	fmt.Println("Server is listening on port 8080")
	http.HandleFunc("/heroes", handlers.HeroesHandler)
	http.HandleFunc("/villains", handlers.VillainsHandler)
	http.ListenAndServe(":8080", nil)
}
