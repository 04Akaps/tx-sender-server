package server

import (
	"github.com/04Akaps/tx-sender-server/config"
	gRPC "github.com/04Akaps/tx-sender-server/gRPC/client"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	config *config.Config
	engine *gin.Engine

	gRPCClient *gRPC.AuthGrpcClient

	port string
}

func NewServer(config *config.Config, gRPCClient *gRPC.AuthGrpcClient) *Server {
	s := &Server{
		config:     config,
		engine:     gin.New(),
		gRPCClient: gRPCClient,
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
