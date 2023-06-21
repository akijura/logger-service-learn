package main

import (
	"context"
	"fmt"
	"log"
	"log-service/logs"
	"net"

	"google.golang.org/grpc"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	//input := req.GetLogEntry()
	res := &logs.LogResponse{
		Result: "here aki",
	}
	return res, nil
}

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("Failed to list for gRPC: %v", err)
	}
	s := grpc.NewServer()

	logs.RegisterLogServiceServer(s, &LogServer{})
	log.Printf("gRPC Server started on port %s", gRpcPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to list for gRPC: %v", err)
	}
}
