CONTRACTS = Migrations SingleNumRegister

CONTRACT_SRCS = $(addprefix contracts/, $(CONTRACTS:=.sol))
CONTRACT_JSONS = $(addprefix build/contracts/, $(CONTRACTS:=.json))
SOLC_ABIS = $(addprefix build/solc_outs/, $(CONTRACTS:=.abi))
SOLC_BINS = $(addprefix build/solc_outs/, $(CONTRACTS:=.bin))

BINDINGS = bindings/SingleNumRegister.go

.PHONY: run gen_bindings

build/contracts/%.json: contracts/%.sol
	truffle migrate

build/solc_outs/%.abi: contracts/%.sol
	solc --optimize --abi $< -o build/solc_outs

build/solc_outs/%.bin: contracts/%.sol
	solc --optimize --bin $< -o build/solc_outs

bindings/%.go: build/solc_outs/%.abi build/solc_outs/%.bin
	abigen --abi=$(word 1,$^) --bin=$(word 2,$^) --pkg=api --out=$@

gen_bindings: $(BINDINGS)

run: $(CONTRACT_JSONS) gen_bindings
	go run github.com/shugo256/golang-with-solidity-poc

clean:
	rm -f $(CONTRACT_JSONS) $(SOLC_ABIS) $(SOLC_BINS) $(BINDINGS)
