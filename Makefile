.PHONY: all server client clean

build:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		pkg/pokemon/pokemon.proto

server:
	go run ./server

client:
	go run ./client
