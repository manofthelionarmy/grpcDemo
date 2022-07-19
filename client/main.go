package main

import (
	"context"
	"fmt"
	pb "grpcDemo/pkg/pokemon"
	"log"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial(
		"localhost:8080",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	grpClient := pb.NewPokemonServiceClient(cc)
	res, err := grpClient.PokeDex(context.Background(), &pb.PokemonRequest{Name: "Pikachu"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}
