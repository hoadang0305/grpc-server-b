package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "github.com/hoadang0305/grpc-server-b/public/controller/http/v1"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	actorHandler *v1.ActorHandler
	filmHandler  *v1.FilmHandler
}

func NewServer(actorHandler *v1.ActorHandler, filmHandler *v1.FilmHandler) *Server {
	return &Server{actorHandler: actorHandler, filmHandler: filmHandler}
}

func (s *Server) Run() {
	router := gin.New()
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	httpServerInstance := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	v1.MapRoutes(router, s.actorHandler, s.filmHandler)
	err := httpServerInstance.ListenAndServe()
	if err != nil {
		return
	}
	fmt.Println("Server running at " + httpServerInstance.Addr)
}
