package main

import (
	"fmt"
	"gophercises/syntax/controller"
	"log"
	"net/http"
)

// var DevMw = devMw

func main() {
	mux := http.NewServeMux()
	fmt.Println("Starting server localhost:3000")

	mux.HandleFunc("/debug/", controller.SourceCodeHandler) //main code with highliting
	mux.HandleFunc("/panic/", controller.PanicDemo)
	mux.HandleFunc("/", controller.Hello)

	log.Fatal(http.ListenAndServe(":3000", controller.DevMw(mux)))
}
