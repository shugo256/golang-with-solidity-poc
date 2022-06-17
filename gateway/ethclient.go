package gateway

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	api "github.com/shugo256/golang-with-solidity-poc/bindings"
	"os"
)

var SingleNumRegister *api.Api

func InitContractClients(url string) error {
	client, err := ethclient.Dial(url)
	if err != nil {
		return err
	}

	SingleNumRegister, err = api.NewApi(common.HexToAddress(os.Getenv("CONTRACT_ADDRESS")), client)

	return err
}
