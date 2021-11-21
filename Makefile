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
%:
	@:
