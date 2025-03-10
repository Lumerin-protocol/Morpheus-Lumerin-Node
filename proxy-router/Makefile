run:
	go run cmd/main.go

run-user:
	WALLET_PRIVATE_KEY=0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a DIAMOND_CONTRACT_ADDRESS=0xDE819AaEE474626E3f34Ef0263373357e5a6C71b MOR_TOKEN_ADDRESS=0x092bAaDB7DEf4C3981454dD9c0A0D7FF07bCFc86 PROXY_STORAGE_PATH='./data/badger2/' PROXY_ADDRESS='0.0.0.0:3334' WEB_ADDRESS='0.0.0.0:8083' BLOCKSCOUT_API_URL=https://arbitrum.blockscan.com/api/v2 ETH_NODE_CHAIN_ID=42161 make run

run-race:
	GOTRACEBACK=crash go run -gcflags '-N -l' -race cmd/main.go

build:
	TAG_NAME=$(TAG_NAME) ./build.sh
	cp .env.example .env 
	cp models-config.json.example models-config.json
	cp rating-config.json.example rating-config.json
	
clean:
	rm -rf bin logs data .cookie proxy.conf proxy-router .env models-config.json rating-config.json

install:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.53.3
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install github.com/praetorian-inc/gokart@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	
format:
	gofmt -s -w .
	
lint:
	golangci-lint run
	govulncheck ./...
	gokart scan .
	gosec ./...

test-unit:
	go test -v -p 1 $$(go list ./... | grep -v /test)

test-integration:
	go test -v ./test/...

swagger:
	swag fmt -g ./internal/handlers/httphandlers/http.go && swag init -g ./internal/handlers/httphandlers/http.go --parseInternal --parseDependency

mocks:
	mockery

reset-keychain:
	go run cmd/main.go -app-reset-keychain

update-bindings:
	cd ../smart-contracts && yarn bindings-go && cp -r bindings/go/contracts/{diamond/facets/*,tokens/*,mock/tokens/*} ../proxy-router/internal/repositories/contracts/bindings
