package server

import (
	"github.com/04Akaps/tx-sender-server/config"
	gRPC "github.com/04Akaps/tx-sender-server/gRPC/client"
	"github.com/04Akaps/tx-sender-server/service"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	config *config.Config
	engine *gin.Engine

	gRPCClient *gRPC.AuthGrpcClient
	service    *service.Service

	port string
}

func NewServer(config *config.Config, gRPCClient *gRPC.AuthGrpcClient, service *service.Service) *Server {
	s := &Server{
		config:     config,
		engine:     gin.New(),
		gRPCClient: gRPCClient,
		service:    service,
		port:       config.Server.Port,
	}

	s.setCors()

	newTx(s)

	return s

}

func (s *Server) StartServer() error {
	log.Printf("Start Tx Server")
	return s.engine.Run(s.port)
}
