AVA_PATH=./internal/avalanche

ava.compile:
	@solc @openzeppelin/=../lock-avax-contracts/node_modules/@openzeppelin/ \
		--abi ../lock-avax-contracts/contracts/*.sol \
		-o $(AVA_PATH)/abi \
		 --overwrite

ava.gen:
	@abigen --abi $(AVA_PATH)/abi/LockManager.abi --pkg=locker --out=$(AVA_PATH)/locker/locker.go

build:
	@go build -o ./build/heimdallr ./cmd/heimdallr/*.go

run:
	@go run ./cmd/heimdallr/*.go

.PHONY: build run