package main

import _ "expvar"
import (
	"http"
	//	"io"
	"log"
)


func subscribe(w http.ResponseWriter, req *http.Request) {

}

func publish(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	log.Println(req.Form)
}

func main() {
	http.HandleFunc("/subscribe", subscribe) // Subsribers must declare themselves using this URI
	http.HandleFunc("/publish", publish)     // producers can publish messages using this URI
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.String())
	}
}

