package app

import (
	"github.com/04Akaps/tx-sender-server/config"
	gRPC "github.com/04Akaps/tx-sender-server/gRPC/client"
	"github.com/04Akaps/tx-sender-server/server"
)

type App struct {
	config     *config.Config
	server     *server.Server
	gRPCClient *gRPC.AuthGrpcClient
}

func NewTxApp(config *config.Config) *App {
	a := &App{
		config: config,
	}

	var err error

	if a.gRPCClient, err = gRPC.NewGrpcClient(config); err != nil {
		panic(err)
	} else {
		a.server = server.NewServer(config, a.gRPCClient)
	}

	return a
}

func (a *App) StartTxApp() error {
	return a.server.StartServer()
}
