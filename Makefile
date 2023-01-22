EVM_PATH=./internal/evm

evm.compile:
	@solc @openzeppelin/=../bifrost-solidity-contracts/node_modules/@openzeppelin/ \
		--abi ../bifrost-solidity-contracts/contracts/*.sol \
		-o $(EVM_PATH)/abi \
		 --overwrite

evm.gen:
	@abigen --abi $(AVA_PATH)/abi/LockManager.abi --pkg=locker --out=$(AVA_PATH)/locker/locker.go

build:
	@go build -o ./build/heimdallr ./cmd/heimdallr/*.go

run:
	@go run ./cmd/heimdallr/*.go

.PHONY: build run