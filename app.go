package main

import (
	"github.com/joho/godotenv"
	"github.com/shugo256/golang-with-solidity-poc/gateway"
	"github.com/shugo256/golang-with-solidity-poc/handlers/single_num"
	"log"
	"net/http"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := gateway.InitContractClients("http://127.0.0.1:7545"); err != nil {
		log.Fatal(err)
	}

	http.Handle("/set/", single_num.Setter)
	http.Handle("/", single_num.Getter)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
