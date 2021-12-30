
test:
	go test ./...

run-wasp:
	wasp -c ./resources/wasp/config.json

wasm:
	make build-wasm
	make deploy-wasm

build-wasm:
	tinygo build -o ./src/iscplotterycontract/test/iscplotterycontract_bg.wasm -target wasm ./src/iscplotterycontract/go/main.go

deploy-wasm:
	wasp-cli chain deploy-contract wasmtime iscplotterycontract "ISCP Lottery" ./src/iscplotterycontract/test/iscplotterycontract_bg.wasm -d --verbose
	
deploy-chain:
	wasp-cli chain deploy --committee=0 --quorum=1 --chain=iscp-lottery-chain --description="ISCP Lottery Chain"
	wasp-cli chain deposit IOTA:10000
