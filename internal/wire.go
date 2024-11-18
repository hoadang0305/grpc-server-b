//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/google/wire"
	"github.com/hoadang0305/grpc-server-b/internal/controller"
	"github.com/hoadang0305/grpc-server-b/internal/controller/grpc"
	filmGrpc "github.com/hoadang0305/grpc-server-b/internal/controller/grpc/v1"
	"github.com/hoadang0305/grpc-server-b/internal/controller/http"
	v1 "github.com/hoadang0305/grpc-server-b/internal/controller/http/v1"
	"github.com/hoadang0305/grpc-server-b/internal/database"
	repositoryimplement "github.com/hoadang0305/grpc-server-b/internal/repository/implement"
	serviceimplement "github.com/hoadang0305/grpc-server-b/internal/service/implement"
)

var container = wire.NewSet(
	controller.NewApiContainer,
)

// may have grpc server in the future
var serverSet = wire.NewSet(
	http.NewServer,
	grpc.NewServer,
)

// handler === controller | with service and repository layers to form 3 layers architecture
var handlerSet = wire.NewSet(
	v1.NewActorHandler,
	v1.NewFilmHandler,
	filmGrpc.NewFilmHandler,
)

var serviceSet = wire.NewSet(
	serviceimplement.NewActorService,
	serviceimplement.NewFilmService,
)

var repositorySet = wire.NewSet(
	repositoryimplement.NewActorRepository,
	repositoryimplement.NewFilmRepository,
)

func InitializeContainer(
	db database.Db,
) *controller.ApiContainer {
	wire.Build(serverSet, handlerSet, serviceSet, repositorySet, container)
	return &controller.ApiContainer{}
}
