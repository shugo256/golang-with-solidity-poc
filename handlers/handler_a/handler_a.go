package handler_a

import (
	"fmt"
	"github.com/shugo256/golang-with-solidity-poc/handlers"
	"net/http"
)

func handlerA(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Hello Solidity!")
	if err != nil {
		return
	}
}

var HandlerA = handlers.HandlerFuncWithOptions(handlerA, handlers.LogWrapper)
