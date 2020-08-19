package main

import (
	"log"
	"net"
	"server/mysql"
	"server/note"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	mysql.SetUp()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	note.RegisterNoteOptionServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
