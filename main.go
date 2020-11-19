package main

import (
	"fmt"
	"net/http"
)

var port string = ":8080"

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/room", HandleRoom)

	http.HandleFunc("/ws", ws)

	fmt.Println("Server starting at " + port)
	http.ListenAndServe(port, nil)
}
