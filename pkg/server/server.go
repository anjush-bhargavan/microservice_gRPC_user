package server

import (
	"log"
	"net"

	"github.com/anjush-bhargavan/microservice_gRPC_user/pkg/handler"
	pb "github.com/anjush-bhargavan/microservice_gRPC_user/pkg/pb"
	"google.golang.org/grpc"
)

func NewGrpcServer(handlr *handler.UserHandler) {
	log.Println("connecting to gRPC server")
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal("error creating listener on port 8082")
	}
	
	grp := grpc.NewServer()
	pb.RegisterUserServiceServer(grp, handlr)

	log.Printf("listening on gRPC server 8082")
	if err := grp.Serve(lis); err != nil {
		log.Fatal("error connecting to gRPC server")

	}
}
