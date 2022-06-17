package single_num

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shugo256/golang-with-solidity-poc/gateway"
	"github.com/shugo256/golang-with-solidity-poc/handlers"
	"math/big"
	"net/http"
	"os"
	"strings"
)

func parseParams(path string) big.Int {
	var n big.Int
	fmt.Println(strings.TrimPrefix(path, "/set/"))
	_, err := fmt.Sscan(strings.TrimPrefix(path, "/set/"), &n)
	if err != nil {
		fmt.Println(err)
		return big.Int{}
	}
	return n
}

func setter(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Setter")
	n := parseParams(req.URL.Path)

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	transaction, err := gateway.SingleNumRegister.Set(bind.NewKeyedTransactor(privateKey), &n)
	if err != nil {
		fmt.Println(err)
	}

	_, err = fmt.Fprintf(w, "Set value: %v (transaction: %v)", n.String(), transaction)
	if err != nil {
		fmt.Println(err)
	}
}

var Setter = handlers.HandlerFuncWithOptions(setter, handlers.LogWrapper)
