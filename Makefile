include .env
export $(shell sed 's/=.*//' .env)

GOPATH=$(shell go env GOPATH)

wq-producer:
	@ echo
	@ echo "Starting producer..."
	@ echo
	@ go run ./cmd/workqueue/producer/main.go

wq-consumer:
	@ echo
	@ echo "Starting consumer..."
	@ echo
	@ go run ./cmd/workqueue/consumer/main.go

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
