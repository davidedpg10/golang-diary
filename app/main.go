package main

import (
	"fmt"
	"net/http"
)

var port string

func init() {
	port = ":8080"

}

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(port, nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world, %s!", r.URL.Path[1:])
}
