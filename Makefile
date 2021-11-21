include .env
export $(shell sed 's/=.*//' .env)

GOPATH=$(shell go env GOPATH)

producer:
	@ echo
	@ echo "Starting producer..."
	@ echo
	@ go run ./cmd/producer/main.go

consumer:
	@ echo
	@ echo "Starting consumer..."
	@ echo
	@ go run ./cmd/consumer/main.go

gen-proto:
	@ echo "Generating proto files..."
	@ protoc \
		--go_out=. \
    --go-grpc_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
    ./proto/*.proto

%:
	@:
