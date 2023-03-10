package server

import (
	"draft/datastore"
	"log"
	"net"

	"draft/grpc"

	ggrpc "google.golang.org/grpc"
)

type Server struct {
	srv *ggrpc.Server
}

type server struct {
	ds *datastore.MemCache
	grpc.UnimplementedBuyersAndShopsServer
}

func NewServer(ds *datastore.MemCache) *Server {
	gs := ggrpc.NewServer()
	grpc.RegisterBuyersAndShopsServer(gs, &server{
		ds: ds,
	})
	srv := &Server{
		srv: gs,
	}
	return srv
}

func (s *Server) Start() {
	defer func() {
		err := recover()
		if err != nil {
			println(err)
			s.Stop()
		}
	}()
	lis, err := net.Listen("tcp", `0.0.0.0:8080`)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := s.srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) Stop() {
}
