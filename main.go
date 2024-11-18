package main

import (
	_ "github.com/hoadang0305/grpc-server-b/docs"
	"github.com/hoadang0305/grpc-server-b/startup"
)

// @title           API for Advanced Web
// @version         1.0
// @description     API for Advanced Web
// @BasePath  /api/v1
func main() {
	startup.Execute()
}
