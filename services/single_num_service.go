package services

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shugo256/golang-with-solidity-poc/gateway"
	singlenumregister "github.com/shugo256/golang-with-solidity-poc/gen/single_num_register"
	"html/template"
	"math/big"
	"os"
)

type SingleNumRegisterService struct{}

func (s SingleNumRegisterService) HTML(ctx context.Context) ([]byte, error) {
	getRes, err := s.GetNum(ctx)
	if err != nil {
		return nil, err
	}

	htmlTemplate, err := template.New("HTML").Parse("<h1>{{.}}</h1>")
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	err = htmlTemplate.Execute(&buf, getRes.Value)
	if err != nil {
		return nil, err
	}

	fmt.Println(buf.String())

	return buf.Bytes(), nil
}

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
