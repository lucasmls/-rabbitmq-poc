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

ps-producer:
	@ echo
	@ echo "Starting PubSub producer..."
	@ echo
	@ go run ./cmd/pubsub/producer/main.go

ps-first-consumer:
	@ echo
	@ echo "Starting PubSub consumer..."
	@ echo
	@ go run ./cmd/pubsub/consumer/main.go

ps-second-consumer:
	@ echo
	@ echo "Starting PubSub consumer..."
	@ echo
	@ go run ./cmd/pubsub/consumer/main.go > consumer_logs.log

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
