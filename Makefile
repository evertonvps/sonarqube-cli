BINARY_NAME=sonarqube-cli
.DEFAULT_GOAL := run

build:
	GOARCH=amd64 GOOS=darwin go build -o ./target/${BINARY_NAME}-darwin ./cmd/...
	GOARCH=amd64 GOOS=linux go build -o ./target/${BINARY_NAME}-linux ./cmd/...
	GOARCH=amd64 GOOS=windows go build -o ./target/${BINARY_NAME}-windows ./cmd/...

run: build
	./target/${BINARY_NAME}-linux

clean:
	go clena  
	rm ./target/${BINARY_NAME}-darwin
	rm ./target/${BINARY_NAME}-linux
	rm ./target/${BINARY_NAME}-windows
	
test:
	go test ./...

test-examples:
	go test --tags=examples ./...

test_coverage:
	go test ./... -coverprofile=coverage.out	

dep:
	go mod download

vet:
	go vet ./...

lint:
	golangci-lint run --enable-all


deploy: build
	mv ./target/${BINARY_NAME}-linux ./target/${BINARY_NAME}
	chmod +x ./target/${BINARY_NAME}
	sudo mv ./target/${BINARY_NAME} /usr/local/bin/${BINARY_NAME}