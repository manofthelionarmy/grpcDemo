package main

import (
	pb "grpcDemo/pkg/pokemon"
	"grpcDemo/pkg/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	pokemonSvc := service.New()

	grpcServer := grpc.NewServer()
	pb.RegisterPokemonServiceServer(grpcServer, pokemonSvc)

	listner, _ := net.Listen("tcp", "localhost:8080")
	err := grpcServer.Serve(listner)
	if err != nil {
		log.Fatal(err)
	}
	defer grpcServer.Stop()
}
