package router

import (
    "log"
    "net"

    "github.com/ARUN-AK5396/demo-grpc/router/Handlers" // Assuming Handlers is the package name
    pb "github.com/ARUN-AK5396/demo-grpc/pb" // Assuming pb is the package name
    "google.golang.org/grpc"
)

func startGRPCServer(listener net.Listener) error {
    grpcServer := grpc.NewServer()
    pb.RegisterRoutesServiceServer(grpcServer, &Handlers.RoutesServer{}) // Using pb and Handlers correctly
    log.Println("Serving gRPC on :8080")
    return grpcServer.Serve(listener)
}
