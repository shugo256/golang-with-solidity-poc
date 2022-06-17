package main

import (
	"github.com/shugo256/golang-with-solidity-poc/handlers/handler_a"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", handler_a.HandlerA)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
