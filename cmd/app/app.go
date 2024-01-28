package app

import (
	"github.com/04Akaps/tx-sender-server/config"
	gRPC "github.com/04Akaps/tx-sender-server/gRPC/client"
	"github.com/04Akaps/tx-sender-server/repository/node"
	"github.com/04Akaps/tx-sender-server/repository/redis"
	"github.com/04Akaps/tx-sender-server/server"
	"github.com/04Akaps/tx-sender-server/service"
)

type App struct {
	config  *config.Config
	server  *server.Server
	service *service.Service

	gRPCClient *gRPC.AuthGrpcClient
	redis      redis.RedisImpl
	node       node.NodeImpl
}

func NewTxApp(config *config.Config) *App {
	a := &App{
		config: config,
	}

	var err error

	if a.gRPCClient, err = gRPC.NewGrpcClient(config); err != nil {
		panic(err)
	} else if a.redis, err = redis.NewRedis(config); err != nil {
		panic(err)
	} else if a.node, err = node.NewNode(config); err != nil {
		panic(err)
	} else {
		a.service = service.NewService(a.node, a.redis)
		a.server = server.NewServer(config, a.gRPCClient, a.service)
	}

	return a
}

func (a *App) StartTxApp() error {
	return a.server.StartServer()
}
