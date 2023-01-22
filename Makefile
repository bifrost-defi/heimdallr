EVM_PATH=./internal/evm

evm.compile:
	@solc @openzeppelin/=../bifrost-solidity-contracts/node_modules/@openzeppelin/ \
		--abi ../bifrost-solidity-contracts/contracts/*.sol \
		-o $(EVM_PATH)/abi \
		 --overwrite

evm.gen:
	@mkdir -p $(EVM_PATH)/wrapping-bridge/
	@abigen --abi $(EVM_PATH)/abi/WrappingBridge.abi --pkg=wrappingBridge --out=$(EVM_PATH)/wrapping-bridge/bridge.go

build:
	@go build -o ./build/heimdallr ./cmd/heimdallr/*.go

run:
	@go run ./cmd/heimdallr/*.go

.PHONY: build run