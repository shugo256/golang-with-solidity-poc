package gateway

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shugo256/golang-with-solidity-poc/bindings"
	"os"
)

var SingleNumRegister *bindings.SingleNumRegister

func InitContractClients(url string) error {
	client, err := ethclient.Dial(url)
	if err != nil {
		return err
	}

	SingleNumRegister, err = bindings.NewSingleNumRegister(common.HexToAddress(os.Getenv("CONTRACT_ADDRESS")), client)

	return err
}
