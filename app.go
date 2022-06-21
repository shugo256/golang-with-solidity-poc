package main

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/shugo256/golang-with-solidity-poc/gateway"
	genlog "github.com/shugo256/golang-with-solidity-poc/gen/log"
	"github.com/shugo256/golang-with-solidity-poc/services"
	"github.com/shugo256/golang-with-solidity-poc/services/single_num_service"
	"net/http"
	"os"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	logger := genlog.New("main", true)

	if err := godotenv.Load(); err != nil {
		logger.Err(err)
	}

	if err := gateway.InitContractClients("http://127.0.0.1:8545"); err != nil {
		logger.Err(err)
	}
	handler := services.NewHttpHandler()
	handler.AddService(single_num_service.New(*handler), "single-num-service")

	logger.Err(http.ListenAndServe(":8080", handler))
}
