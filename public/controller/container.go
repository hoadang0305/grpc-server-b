package controller

import (
	"github.com/hoadang0305/grpc-server-b/public/controller/grpc"
	"github.com/hoadang0305/grpc-server-b/public/controller/http"
)

type ApiContainer struct {
	HttpServer *http.Server
	GrpcServer *grpc.Server
}

func NewApiContainer(httpServer *http.Server, grpcServer *grpc.Server) *ApiContainer {
	return &ApiContainer{HttpServer: httpServer, GrpcServer: grpcServer}
}
