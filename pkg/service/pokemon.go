package service

import (
	"context"
	"fmt"
	pb "grpcDemo/pkg/pokemon"
)

var _ pb.PokemonServiceServer = (*Pokemon)(nil)

// Data represents the stats and name of our Pokemon
type Data struct {
	Name      string
	Attack    int32
	Hp        int32
	Defense   int32
	SpDefense int32
	SpAttack  int32
	Speed     int32
}

// Pokemon is our service
type Pokemon struct {
	pb.UnimplementedPokemonServiceServer
	PokemonDb map[string]*Data
}

var pokemon = make([]*Data, 0)

// New creates a new Pokemon gRPC service
func New() *Pokemon {
	return &Pokemon{}
}

// PokeDex returns the a pokemon
func (pdx *Pokemon) PokeDex(ctx context.Context, request *pb.PokemonRequest) (*pb.PokemonResponse, error) {
	if v, ok := pdx.PokemonDb[request.Name]; ok {
		return &pb.PokemonResponse{
			Name:      v.Name,
			Hp:        v.Hp,
			Atttak:    v.Attack,
			SpAttack:  v.SpAttack,
			SpDefense: v.SpDefense,
			Speed:     v.Speed,
		}, nil
	}
	return nil, fmt.Errorf("pokemon %s not found", request.Name)
}

func init() {
	pokemon = append(
		pokemon,
		[]*Data{
			{
				Name:      "Pikachu",
				Attack:    int32(55),
				Defense:   int32(45),
				SpAttack:  int32(55),
				SpDefense: int32(40),
				Speed:     int32(65),
			},
		}...,
	)
}
