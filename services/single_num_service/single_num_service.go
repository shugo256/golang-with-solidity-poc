package single_num_service

import (
	"bytes"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shugo256/golang-with-solidity-poc/gateway"
	"github.com/shugo256/golang-with-solidity-poc/gen/http/single_num_register/server"
	singlenumregister "github.com/shugo256/golang-with-solidity-poc/gen/single_num_register"
	"github.com/shugo256/golang-with-solidity-poc/services"
	goahttp "goa.design/goa/v3/http"
	"html/template"
	"math/big"
	"os"
)

type singleNumService struct{}

func New(handler services.HttpHandler) *server.Server {
	endpoints := singlenumregister.NewEndpoints(singleNumService{})

	return server.New(
		endpoints,
		handler,
		goahttp.RequestDecoder,
		goahttp.ResponseEncoder,
		nil,
		nil)
}

func (s singleNumService) HTML(ctx context.Context) ([]byte, error) {
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

	return buf.Bytes(), nil
}

func (s singleNumService) GetNum(ctx context.Context) (*singlenumregister.GetResult, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	value, err := gateway.SingleNumRegister.Get(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	return &singlenumregister.GetResult{Value: singlenumregister.Value(value.Int64())}, nil
}

func (s singleNumService) SetNum(ctx context.Context, payload *singlenumregister.SetNumPayload) (*singlenumregister.SetResult, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	privateKey, _ := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	_, err := gateway.SingleNumRegister.Set(bind.NewKeyedTransactor(privateKey), big.NewInt(int64(payload.Val)))
	if err != nil {
		return nil, err
	}

	return &singlenumregister.SetResult{Success: true}, nil
}
