package cmd

import (
	"ewallet-ums/helpers"
	"log"
	"net"

	"google.golang.org/grpc"
)

func ServerGRPC() {
	lis, err := net.Listen("tcp", ":" + helpers.GetEnv("GRPC_PORT","7001"))
	if err != nil {
		log.Fatal("failed to listen grpc port: ", err)
	}
	s := grpc.NewServer()
	// list method
	// pb.ExampleMethod(s, &service.WalletService{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve grpc port: ", err)
	}
}