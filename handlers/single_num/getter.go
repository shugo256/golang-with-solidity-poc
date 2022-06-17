package single_num

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/shugo256/golang-with-solidity-poc/gateway"
	"github.com/shugo256/golang-with-solidity-poc/handlers"
	"net/http"
)

func getter(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("Getter")
	value, err := gateway.SingleNumRegister.Get(&bind.CallOpts{})
	if err != nil {
		return
	}
	_, err = fmt.Fprintf(w, value.String())
	if err != nil {
		return
	}
}

var Getter = handlers.HandlerFuncWithOptions(getter, handlers.LogWrapper)
