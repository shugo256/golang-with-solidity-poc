CONTRACTS = Migrations SingleNumRegister

CONTRACT_SRCS = $(addprefix contracts/, $(CONTRACTS:=.sol))
CONTRACT_JSONS = $(addprefix build/contracts/, $(CONTRACTS:=.json))
SOLC_ABIS = $(addprefix build/solc_outs/, $(CONTRACTS:=.abi))
SOLC_BINS = $(addprefix build/solc_outs/, $(CONTRACTS:=.bin))

BINDINGS = bindings/SingleNumRegister.go

GOA_DESIGN_DIR = design
GOA_GEN_DIR = gen

.PHONY: run bindings goa

build/contracts/%.json: contracts/%.sol
	truffle migrate

build/solc_outs/%.abi: contracts/%.sol
	solc --optimize --abi $< -o build/solc_outs

build/solc_outs/%.bin: contracts/%.sol
	solc --optimize --bin $< -o build/solc_outs

bindings/%.go: build/solc_outs/%.abi build/solc_outs/%.bin
	abigen --abi=$(word 1,$^) --bin=$(word 2,$^) --pkg=bindings --type $* --out $@

$(GOA_GEN_DIR): $(GOA_DESIGN_DIR)
	goa gen github.com/shugo256/golang-with-solidity-poc/$(GOA_DESIGN_DIR)

bindings: $(BINDINGS)

goa: $(GOA_GEN_DIR)

run: $(CONTRACT_JSONS) bindings goa
	go run github.com/shugo256/golang-with-solidity-poc

clean:
	rm -rf $(CONTRACT_JSONS) $(SOLC_ABIS) $(SOLC_BINS) $(BINDINGS) $(GOA_GEN_DIR)
