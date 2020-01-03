package main

import (
	"gophercises/transform/controller"
	"log"
	"net/http"
)

func main() {
	// go func() {
	// 	t := time.NewTicker(1 * time.Minute)
	// 	for {
	// 		<-t.C
	// 	}
	// }()
	fs := http.FileServer(http.Dir("./img/"))

	mux := http.NewServeMux()
	mux.HandleFunc("/", controller.Homepage)
	mux.HandleFunc("/modify/", controller.Modify)
	mux.HandleFunc("/upload", controller.Upload)
	mux.Handle("/img/", http.StripPrefix("/img", fs))

	//mux.HandleFunc("/modify/{someParameter}", controller.Modify)
	//this is optional url because this is main and not testing environment
	//because image name and modes will append inside the {} upon calling
	//but in main this append is handled in the corrosponding func

	log.Fatal(http.ListenAndServe(":3000", mux))
}
