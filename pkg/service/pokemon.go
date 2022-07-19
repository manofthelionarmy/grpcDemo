package service

import (
	"context"
	pb "grpcDemo/pkg/pokemon"
)

var _ pb.PokemonServiceServer = (*Pokemon)(nil)

// Data represents the stats and name of our Pokemon
type Data struct {
	Name      string
	Attack    int
	Defense   int
	SpDefense int
	SpAttack  int
	Speed     int
}

// Pokemon is our service
type Pokemon struct {
	pb.UnimplementedPokemonServiceServer
}

// PokeDex returns the a pokemon
func (pdx *Pokemon) PokeDex(ctx context.Context, request *pb.PokemonRequest) (*pb.PokemonResponse, error) {
	return nil, nil
}
