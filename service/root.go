package service

import (
	"github.com/04Akaps/tx-sender-server/repository/node"
	"github.com/04Akaps/tx-sender-server/repository/redis"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"log"
)

type Service struct {
	node  node.NodeImpl
	redis redis.RedisImpl
}

func NewService(node node.NodeImpl, redis redis.RedisImpl) *Service {
	return &Service{node, redis}
}

func (s *Service) GetABI(bytes []byte) (*abi.ABI, error) {
	//type ABI struct {
	//	Functions []Function `json:"functions"`
	//}
	//
	//type Function struct {
	//	Name       string `json:"name"`
	//	InputType  string `json:"input"`
	//	OutputType string `json:"output"`
	//}

	//var contractABI ABI
	//err = json.Unmarshal(abiData, &contractABI)

	// TODO how to get ABI From Bytes??

	return nil, nil
}

func (s *Service) GetCode(chain, address string) ([]byte, error) {

	//isContractAddress := len(byteCode) > 0

	if res, err := s.node.GetCodeAt(chain, address); err != nil {
		log.Println("Failed To Get Code", "chain", chain, "address", address)
		return nil, err
	} else {
		return res, nil
	}
}
