package server

import (
	"fmt"
	"log"

	"github.com/mainnika/private-tv/rc"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// GetState GetState
func (_ Server) GetState(ctx context.Context, in *rc.Empty) (*rc.State, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("no context")
	}

	log.Println(md)

	return &rc.State{Source: "foobar"}, nil
}
