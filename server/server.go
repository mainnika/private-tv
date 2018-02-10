package server

import (
	"fmt"
	"log"
	"net"

	"github.com/mainnika/private-tv/rc"
	"google.golang.org/grpc"
)

// Server Server
type Server struct {
	net.Listener
	grpc *grpc.Server
}

// NewServer NewServer
func NewServer(port int) (err error, s *Server) {

	s = &Server{grpc: grpc.NewServer()}

	s.Listener, err = net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	rc.RegisterPrivateTVServer(s.grpc, s)

	return nil, s
}

// Serve Serve
func (s *Server) Serve() {

	s.grpc.Serve(s)
}
