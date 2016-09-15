//testvpp 4.3
package main

import (
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	log.Println("starting web server...")

	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8080", nil)
}
