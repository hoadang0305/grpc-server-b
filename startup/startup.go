package startup

import (
	"github.com/gammazero/workerpool"
	"github.com/hoadang0305/grpc-server-b/public"
	"github.com/hoadang0305/grpc-server-b/public/controller"
	"github.com/hoadang0305/grpc-server-b/public/database"
	"github.com/hoadang0305/grpc-server-b/public/utils/validation"
)

func registerDependencies() *controller.ApiContainer {
	// Open database connection
	db := database.Open()

	return public.InitializeContainer(db)
}

func startServers(container *controller.ApiContainer) {
	wp := workerpool.New(2)

	wp.Submit(container.HttpServer.Run)
	wp.Submit(container.GrpcServer.Run)

	wp.StopWait()
}

func Execute() {
	validation.GetValidations()
	container := registerDependencies()
	startServers(container)
}
