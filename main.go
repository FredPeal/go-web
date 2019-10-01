package main

import(
	"net/http"
	"log"
)

func main() {
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))
}