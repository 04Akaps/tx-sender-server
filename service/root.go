package service

import (
	"github.com/04Akaps/tx-sender-server/repository/node"
	"github.com/04Akaps/tx-sender-server/repository/redis"
)

type Service struct {
	node  *node.NodeClient
	redis redis.RedisImpl
}

func NewService(node *node.NodeClient, redis redis.RedisImpl) *Service {
	return &Service{node, redis}
}
