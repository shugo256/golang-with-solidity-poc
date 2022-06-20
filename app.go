package main

import (
	"github.com/joho/godotenv"
	"github.com/shugo256/golang-with-solidity-poc/gateway"
	"github.com/shugo256/golang-with-solidity-poc/gen/http/single_num_register/server"
	singlenumregister "github.com/shugo256/golang-with-solidity-poc/gen/single_num_register"
	"github.com/shugo256/golang-with-solidity-poc/services"
	goahttp "goa.design/goa/v3/http"
	"log"
	"net/http"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := gateway.InitContractClients("http://127.0.0.1:8545"); err != nil {
		log.Fatal(err)
	}

	endpoints := singlenumregister.NewEndpoints(services.SingleNumRegisterService{})
	handler := goahttp.NewMuxer()
	server.Mount(handler, server.New(endpoints, handler, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil))

	log.Fatal(http.ListenAndServe(":8080", handler))
}
