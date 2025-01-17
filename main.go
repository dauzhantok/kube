package kube

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func sayhello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi!")
}
