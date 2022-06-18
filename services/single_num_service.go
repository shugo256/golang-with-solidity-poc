package services

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shugo256/golang-with-solidity-poc/gateway"
	singlenumregister "github.com/shugo256/golang-with-solidity-poc/gen/single_num_register"
	"math/big"
	"os"
)

type SingleNumRegisterService struct{}

func (s SingleNumRegisterService) GetNum(ctx context.Context) (*singlenumregister.GetResult, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	fmt.Printf("Setter: %v\n", ctx.Err())

	value, err := gateway.SingleNumRegister.Get(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	return &singlenumregister.GetResult{Value: singlenumregister.Value(value.Int64())}, nil
}

func (s SingleNumRegisterService) SetNum(ctx context.Context, payload *singlenumregister.SetNumPayload) (*singlenumregister.SetResult, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	fmt.Printf("Setter: %v\n", ctx.Err())

	privateKey, _ := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	transaction, err := gateway.SingleNumRegister.Set(bind.NewKeyedTransactor(privateKey), big.NewInt(int64(payload.Val)))
	if err != nil {
		return nil, err
	}
	fmt.Println(transaction)

	return &singlenumregister.SetResult{Success: true}, nil
}
